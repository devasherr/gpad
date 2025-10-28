package parser

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
)

var typeInfo = map[string]struct {
	size, align int
}{
	"bool":       {1, 1},
	"int8":       {1, 1},
	"uint8":      {1, 1},
	"int16":      {2, 2},
	"uint16":     {2, 2},
	"int32":      {4, 4},
	"uint32":     {4, 4},
	"float32":    {4, 4},
	"int64":      {8, 8},
	"uint64":     {8, 8},
	"float64":    {8, 8},
	"complex64":  {8, 8},
	"complex128": {16, 8},
	"string":     {16, 8},
	"uintptr":    {8, 8},
	"int":        {8, 8},
	"uint":       {8, 8},
}

func typeSize(expr ast.Expr) int {
	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "bool":
			return 0
		case "int8", "uint8", "byte":
			return 1
		case "int16", "uint16":
			return 2
		case "int32", "uint32", "float32", "rune":
			return 4
		case "int64", "uint64", "float64", "int", "uint", "uintptr":
			return 8
		case "complex64":
			return 8
		case "complex128":
			return 16
		case "string":
			return 16
		default:
			return 8
		}
	default:
		return 8
	}
}

func sortFieldList(fields []*ast.Field) []*ast.Field {
	sort.Slice(fields, func(i, j int) bool {
		return typeSize(fields[i].Type) > typeSize(fields[j].Type)
	})

	return fields
}

func calculateStructMemory(fields []*ast.Field) int {
	offset, maxAlign := 0, 0

	for _, field := range fields {
		info := typeInfo[fmt.Sprintf("%s", field.Type)]
		padding := (info.align - (offset % info.align)) % info.align
		offset += padding + info.size

		if info.align > maxAlign {
			maxAlign = info.align
		}
	}

	totalSize := (offset + maxAlign - 1) / maxAlign * maxAlign
	return totalSize
}

func parseFile(path string, verbose bool) (int, int) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Printf("failed to parse file %s\n", path)
		return -1, -1
	}

	beforeSize, afterSize := 0, 0

	ast.Inspect(file, func(n ast.Node) bool {
		s, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if st, ok := s.Type.(*ast.StructType); ok {
			if verbose {
				beforeSize += calculateStructMemory(st.Fields.List)
			}

			orderedFieldList := sortFieldList(st.Fields.List)

			if verbose {
				afterSize += calculateStructMemory(orderedFieldList)
			}

			st.Fields.List = orderedFieldList
		}

		return true
	})

	f, err := os.OpenFile(path, os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("failed to open file %s\n", path)
	}

	format.Node(f, fset, file)
	return beforeSize, afterSize
}

func ParseFiles(filePaths []string, verbose bool) (int, int) {
	totalBefore, totalAfter := 0, 0
	for _, filePath := range filePaths {
		before, after := parseFile(filePath, verbose)

		totalBefore += before
		totalAfter += after
	}
	return totalBefore, totalAfter
}

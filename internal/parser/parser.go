package parser

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
)

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

func parseFile(path string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Printf("failed to parse file %s\n", path)
		return
	}

	ast.Inspect(file, func(n ast.Node) bool {
		s, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if st, ok := s.Type.(*ast.StructType); ok {
			orderedFieldList := sortFieldList(st.Fields.List)
			st.Fields.List = orderedFieldList
		}

		return true
	})

	f, err := os.OpenFile(path, os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("failed to open file %s\n", path)
	}

	format.Node(f, fset, file)
}

func ParseFiles(filePaths []string) {
	for _, filePath := range filePaths {
		parseFile(filePath)
	}
}

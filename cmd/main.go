package main

import (
	"flag"
	"log"
	"os"

	"github.com/devasherr/gpad/internal/files"
	"github.com/devasherr/gpad/internal/parser"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	input := flag.String("path", "", "file path")
	flag.Parse()

	filePaths := files.CollectFiles(path +"/"+ *input)
	parser.ParseFiles(filePaths)
}

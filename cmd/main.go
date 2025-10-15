package main

import (
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

	filePaths := files.CollectFiles(path + "/test")
	parser.ParseFiles(filePaths)
}

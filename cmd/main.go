package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devasherr/gpad/internal/files"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filesPath := files.CollectFiles(path + "/test")
	fmt.Println(filesPath)
}

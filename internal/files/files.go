package files

import (
	"os"
	"path/filepath"
)

func CollectFiles(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	files := []string{}

	for _, entry := range entries {
		cur_path := path + "/" + entry.Name()

		if entry.IsDir() {
			files = append(files, CollectFiles(cur_path)...)
		} else {
			ext := filepath.Ext(entry.Name())
			if ext == ".go" {
				files = append(files, path+"/"+entry.Name())
			}
		}
	}

	return files
}

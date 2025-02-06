package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "."
	}

	files := listFiles(dir)

	for _, v := range files {
		println(v)
	}
}

func listFiles(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, v := range entries {
		if v.IsDir() {
			continue
		}
		if filepath.Ext(v.Name()) != ".md" {
			continue
		}
		files = append(files, filepath.Join(dir, v.Name()))
	}

	return files
}

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// go run sumdir.go -dir ~ -hidden

func listFilesByExtension(directory string, showHidden bool) (map[string][]string, error) {
	fileGroups := make(map[string][]string)

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !showHidden && file.Name()[0] == '.' {
			continue // Skip hidden files
		}

		if file.IsDir() {
			continue
		}

		extension := filepath.Ext(file.Name())
		fileGroups[extension] = append(fileGroups[extension], file.Name())
	}

	// Sort each group of files
	for _, files := range fileGroups {
		sort.Strings(files)
	}

	return fileGroups, nil
}

func main() {
	// Define command-line flags
	directory := flag.String("dir", ".", "The directory to list files from")
	showHidden := flag.Bool("hidden", false, "Include hidden files in the listing")

	// Parse command-line flags
	flag.Parse()

	fileGroups, err := listFilesByExtension(*directory, *showHidden)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Sort the keys (file extensions) for alphabetical order
	var keys []string
	for key := range fileGroups {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print the grouped and sorted files
	for _, extension := range keys {
		files := fileGroups[extension]
		fmt.Printf("%s:\n", extension)
		for _, file := range files {
			fmt.Printf("%s\n", file)
		}
		fmt.Println()
	}
}

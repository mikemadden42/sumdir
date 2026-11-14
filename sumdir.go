// Command sumdir lists the files in a directory grouped by file extension.
package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// go run sumdir.go -dir ~ -hidden

// noExtension is the group name for files without an extension.
const noExtension = "(none)"

// fileExtension returns the lowercased extension of name without the leading
// dot. A leading dot (as in ".bashrc") marks a hidden file, not an extension.
func fileExtension(name string) string {
	ext := filepath.Ext(strings.TrimPrefix(name, "."))
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	if ext == "" {
		return noExtension
	}
	return ext
}

func listFilesByExtension(directory string, showHidden bool) (map[string][]string, error) {
	fileGroups := make(map[string][]string)

	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		name := entry.Name()
		if !showHidden && strings.HasPrefix(name, ".") {
			continue // Skip hidden files
		}

		isDir := entry.IsDir()
		if entry.Type()&os.ModeSymlink != 0 {
			// Follow symlinks so links to directories are skipped too.
			// Broken links are kept and listed as files.
			if info, err := os.Stat(filepath.Join(directory, name)); err == nil {
				isDir = info.IsDir()
			}
		}
		if isDir {
			continue
		}

		extension := fileExtension(name)
		fileGroups[extension] = append(fileGroups[extension], name)
	}

	// Sort each group of files
	for _, names := range fileGroups {
		slices.Sort(names)
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
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Print the groups in alphabetical order of extension
	for _, extension := range slices.Sorted(maps.Keys(fileGroups)) {
		fmt.Printf("%s:\n", extension)
		for _, name := range fileGroups[extension] {
			fmt.Printf("- %s\n", name)
		}
		fmt.Println()
	}
}

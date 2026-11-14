package main

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestFileExtension(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"main.go", "go"},
		{"MAIN.GO", "go"},
		{"archive.tar.gz", "gz"},
		{"Makefile", noExtension},
		{".bashrc", noExtension},
		{".env.local", "local"},
		{"trailing.", noExtension},
	}
	for _, tt := range tests {
		if got := fileExtension(tt.name); got != tt.want {
			t.Errorf("fileExtension(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestListFilesByExtension(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"b.go", "a.go", "C.GO", "README.md", "Makefile", ".bashrc"} {
		if err := os.WriteFile(filepath.Join(dir, name), nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Mkdir(filepath.Join(dir, "sub.d"), 0o750); err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS != "windows" {
		if err := os.Symlink(filepath.Join(dir, "sub.d"), filepath.Join(dir, "link.d")); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		name       string
		showHidden bool
		want       map[string][]string
	}{
		{
			name: "visible only",
			want: map[string][]string{
				"go":        {"C.GO", "a.go", "b.go"},
				"md":        {"README.md"},
				noExtension: {"Makefile"},
			},
		},
		{
			name:       "with hidden",
			showHidden: true,
			want: map[string][]string{
				"go":        {"C.GO", "a.go", "b.go"},
				"md":        {"README.md"},
				noExtension: {".bashrc", "Makefile"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := listFilesByExtension(dir, tt.showHidden)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListFilesByExtensionMissingDir(t *testing.T) {
	if _, err := listFilesByExtension(filepath.Join(t.TempDir(), "missing"), false); err == nil {
		t.Error("expected an error for a missing directory")
	}
}

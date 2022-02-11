package md

import (
	"os"
	"path/filepath"
	"strings"
)

// IsMarkdownFile returns true if the file is a markdown file.
func IsMarkdownFile(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	}

	if fi.IsDir() {
		return false
	}

	filename := fi.Name()
	ext := filepath.Ext(filename)

	if strings.HasPrefix(filename, ".") {
		return false
	}

	if ext != "md" && ext != "markdown" {
		return false
	}

	return true
}

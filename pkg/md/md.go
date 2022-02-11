package md

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/88250/lute"
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

	if ext != ".md" && ext != ".markdown" {
		return false
	}

	return true
}

// FormatMarkdown formats markdown from reader to writer.
func FormatMarkdown(in io.Reader, out io.Writer) error {
	// Read
	markdownContent, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	// Format
	luteEngine := lute.New()
	formatContent := luteEngine.FormatStr("md", string(markdownContent))

	// Output
	_, err = out.Write([]byte(formatContent))
	if err != nil {
		return err
	}

	return nil
}

// ProcessMDFile formats markdown file.
func ProcessMDFile(filePath string, write bool) error {
	// Get Reader
	in, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer in.Close()

	// Get Writer
	out := os.Stdout
	if write {
		out, err = os.OpenFile(filePath, os.O_WRONLY, 0o666)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	// Format & Rewrite
	err = FormatMarkdown(in, out)
	if err != nil {
		return err
	}

	return nil
}

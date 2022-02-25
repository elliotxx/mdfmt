package md

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/88250/lute"
	"github.com/elliotxx/mdfmt/pkg/merrors"
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
func FormatMarkdown(in io.Reader, out io.Writer) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = merrors.NewCrash(fmt.Errorf("%v", r), string(debug.Stack()))
		}
	}()

	// Read
	markdownContent, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	// Format
	luteEngine := lute.New(func(l *lute.Lute) {
		l.RenderOptions.AutoSpace = true
	})
	formatContent := luteEngine.Format("md", markdownContent)

	// Output
	_, err = out.Write(formatContent)
	if err != nil {
		return err
	}

	return nil
}

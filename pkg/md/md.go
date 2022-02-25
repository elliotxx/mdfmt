package md

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/88250/lute"
	"github.com/elliotxx/mdfmt/pkg/merrors"
	"github.com/pkg/diff"
	"github.com/pkg/diff/write"
)

// fdSem guards the number of concurrently-open file descriptors.
//
// For now, this is arbitrarily set to 200, based on the observation that many
// platforms default to a kernel limit of 256. Ideally, perhaps we should derive
// it from rlimit on platforms that support that system call.
//
// File descriptors opened from outside of this package are not tracked,
// so this limit may be approximate.
var fdSem = make(chan bool, 200)

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

// ProcessMDFile formats markdown file.
func ProcessMDFile(filePath string, write, diff, list bool) error {
	// Get Reader and Source
	source, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	in := bytes.NewReader(source)

	// Get Writer and Target
	out := bytes.NewBuffer([]byte{})

	err = FormatMarkdown(in, out)
	if err != nil {
		return err
	}

	target := out.Bytes()

	if !bytes.Equal(source, target) {
		// List
		if list {
			os.Stdout.WriteString(filePath + "\n")
		}
		// Write
		if write {
			info, err := os.Stat(filePath)
			if err != nil {
				return err
			}

			fdSem <- true

			err = os.WriteFile(filePath, target, info.Mode().Perm())

			<-fdSem

			if err != nil {
				return err
			}
		}

		// Diff
		if diff {
			data, err := diffWithReplaceTempFile(source, target, filePath)
			if err != nil {
				return fmt.Errorf("computing diff: %s", err)
			}

			fmt.Fprintf(os.Stdout, "diff -u %s %s\n", filepath.ToSlash(filePath+".orig"), filepath.ToSlash(filePath))
			os.Stdout.Write(data)
		}
	}

	if !write && !diff && !list {
		_, err = os.Stdout.Write(target)
	}

	return err
}

func diffWithReplaceTempFile(b1, b2 []byte, filename string) ([]byte, error) {
	data := bytes.NewBufferString("")
	err := diff.Text("origin", "target", b1, b2, data, write.TerminalColor())

	if len(data.Bytes()) > 0 {
		return replaceTempFilename(data.Bytes(), filename)
	}

	return data.Bytes(), err
}

// replaceTempFilename replaces temporary filenames in diff with actual one.
//
// --- /tmp/gofmt316145376	2017-02-03 19:13:00.280468375 -0500
// +++ /tmp/gofmt617882815	2017-02-03 19:13:00.280468375 -0500
// ...
// ->
// --- path/to/file.md.orig	2017-02-03 19:13:00.280468375 -0500
// +++ path/to/file.md	2017-02-03 19:13:00.280468375 -0500
// ...
func replaceTempFilename(diff []byte, filename string) ([]byte, error) {
	bs := bytes.SplitN(diff, []byte{'\n'}, 3)
	if len(bs) < 3 {
		return nil, fmt.Errorf("got unexpected diff for %s", filename)
	}
	// Preserve timestamps.
	var t0, t1 []byte
	if i := bytes.LastIndexByte(bs[0], '\t'); i != -1 {
		t0 = bs[0][i:]
	}

	if i := bytes.LastIndexByte(bs[1], '\t'); i != -1 {
		t1 = bs[1][i:]
	}

	// Always print filepath with slash separator.
	f := filepath.ToSlash(filename)
	bs[0] = []byte(fmt.Sprintf("--- %s%s", f+".orig", t0))
	bs[1] = []byte(fmt.Sprintf("+++ %s%s", f, t1))

	return bytes.Join(bs, []byte{'\n'}), nil
}

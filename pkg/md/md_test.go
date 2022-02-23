package md

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsMarkdownFile(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is-markdown-file-1",
			args: args{
				filename: "./testdata/hello.md",
			},
			want: true,
		},
		{
			name: "is-markdown-file-2",
			args: args{
				filename: "./testdata/hello.markdown",
			},
			want: true,
		},
		{
			name: "is-not-markdown-file-1",
			args: args{
				filename: "./testdata/.hello.md",
			},
			want: false,
		},
		{
			name: "is-not-markdown-file-2",
			args: args{
				filename: "./testdata/hello.txt",
			},
			want: false,
		},
		{
			name: "fail-as-not-exist",
			args: args{
				filename: "./testdata/hello-fail.txt",
			},
			want: false,
		},
		{
			name: "fail-as-dir",
			args: args{
				filename: "./testdata/",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMarkdownFile(tt.args.filename); got != tt.want {
				t.Errorf("IsMarkdownFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatMarkdown(t *testing.T) {
	type args struct {
		in io.Reader
	}

	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				in: strings.NewReader("# Hello World"),
			},
			wantOut: "# Hello World\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := FormatMarkdown(tt.args.in, out); (err != nil) != tt.wantErr {
				t.Errorf("FormatMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("FormatMarkdown() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestProcessMDFile(t *testing.T) {
	mockMDFile := "./testdata/hello-temp.md"
	_ = os.WriteFile(mockMDFile, []byte("# Hello World"), 0644)

	defer os.Remove(mockMDFile)

	type args struct {
		filePath string
		write    bool
		diff     bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy-path",
			args: args{
				filePath: "./testdata/hello.md",
				write:    false,
				diff:     false,
			},
			wantErr: false,
		},
		{
			name: "md-file-not-exist",
			args: args{
				filePath: "./testdata/hello-not-exist.md",
				write:    true,
				diff:     false,
			},
			wantErr: true,
		},
		{
			name: "diff",
			args: args{
				filePath: "./testdata/hello.md",
				write:    false,
				diff:     true,
			},
			wantErr: false,
		},
		{
			name: "write",
			args: args{
				filePath: mockMDFile,
				write:    true,
				diff:     false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessMDFile(tt.args.filePath, tt.args.write, tt.args.diff); (err != nil) != tt.wantErr {
				t.Errorf("ProcessMDFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

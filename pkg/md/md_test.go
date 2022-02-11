package md

import "testing"

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
			want: false,
		},
		{
			name: "is-markdown-file-2",
			args: args{
				filename: "./testdata/hello.markdown",
			},
			want: false,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMarkdownFile(tt.args.filename); got != tt.want {
				t.Errorf("IsMarkdownFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

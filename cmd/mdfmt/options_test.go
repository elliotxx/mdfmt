package main

import (
	"os"
	"testing"
)

func TestProcessMDFile(t *testing.T) {
	mockMDFile := "./testdata/hello-temp.md"
	_ = os.WriteFile(mockMDFile, []byte("# Hello World"), 0644)

	defer os.Remove(mockMDFile)

	type args struct {
		filePath string
		write    bool
		diff     bool
		list     bool
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
				list:     false,
			},
			wantErr: false,
		},
		{
			name: "md-file-not-exist",
			args: args{
				filePath: "./testdata/hello-not-exist.md",
				write:    true,
				diff:     false,
				list:     false,
			},
			wantErr: true,
		},
		{
			name: "diff",
			args: args{
				filePath: "./testdata/hello.md",
				write:    false,
				diff:     true,
				list:     false,
			},
			wantErr: false,
		},
		{
			name: "write",
			args: args{
				filePath: mockMDFile,
				write:    true,
				diff:     false,
				list:     false,
			},
			wantErr: false,
		},
		{
			name: "list",
			args: args{
				filePath: "./testdata/hello.md",
				write:    false,
				diff:     false,
				list:     true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessMDFile(tt.args.filePath, tt.args.write, tt.args.diff, tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("ProcessMDFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

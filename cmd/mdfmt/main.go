package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/elliotxx/mdfmt/pkg/md"
	"github.com/elliotxx/mdfmt/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cmdShort   = i18n.T(`A Markdown formatter that follow the CommonMark`)
	cmdLong    = templates.LongDesc(i18n.T(`A Markdown formatter that follow the CommonMark. Like gofmt, but for Markdown.`))
	cmdExample = templates.Examples(i18n.T(`
		# Format specified Markdown file, and write to stdout
		mdfmt README.md

		# Format and rewrite for specified Markdown file
		mdfmt -w README.md

		# Display diffs instead of rewriting Markdown files
		mdfmt -d README.md

		# List files whose formatting differs from mdfmt's
		mdfmt -l .

		# Format, rewrite, and display diffs for specified Markdown file
		mdfmt -d -w README.md

		# Format and rewrite all Markdown file in current directory
		mdfmt -w *.md

		# Recursive format and rewrite all Markdown file in current directory
		mdfmt -w .

		# Format and rewrite the specified Markdown file and directory
		mdfmt -w README.md testdir/

		# Format stdin to stdout
		cat README.md | mdfmt

		# Show version info
		mdfmt -V
		`))
)

type Options struct {
	ShowVersion bool
	Write       bool
	Diff        bool
	List        bool
}

// NewEncryptOptions returns a new EncryptOptions instance
func NewMainOptions() *Options {
	return &Options{}
}

func configureCLI() *cobra.Command {
	o := NewMainOptions()
	rootCmd := &cobra.Command{
		Use:          "mdfmt [flags] [path ...]",
		Short:        cmdShort,
		Long:         cmdLong,
		Example:      cmdExample,
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, args []string) (err error) {
			// Show version info
			if o.ShowVersion {
				fmt.Println(version.YAML())
				return nil
			}

			// Process input
			if len(args) == 0 {
				// Stdin
				err = md.FormatMarkdown(os.Stdin, os.Stdout)
				return err
			}
			for _, p := range args {
				// File or directory
				err = filepath.WalkDir(p, func(path string, d fs.DirEntry, _ error) error {
					if !d.IsDir() && md.IsMarkdownFile(path) {
						return md.ProcessMDFile(path, o.Write, o.Diff, o.List)
					}
					return nil
				})
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	rootCmd.Flags().BoolVarP(&o.ShowVersion, "version", "V", false, "show version info")
	rootCmd.Flags().BoolVarP(&o.Write, "write", "w", false, "write result to (source) file instead of stdout")
	rootCmd.Flags().BoolVarP(&o.Diff, "diff", "d", false, "display diffs instead of rewriting files")
	rootCmd.Flags().BoolVarP(&o.List, "list", "l", false, "list files whose formatting differs from mdfmt's")

	return rootCmd
}

func main() {
	// init rootCmd
	rootCmd := configureCLI()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

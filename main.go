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
		# Format markdown file, and write to stdout
		mdfmt README.md

		# Format and rewrite markdown file
		mdfmt -w README.md

		# Format and rewrite markdown file and directory
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
				return md.FormatMarkdown(os.Stdin, os.Stdout)
			}
			for _, p := range args {
				err = filepath.WalkDir(p, func(path string, d fs.DirEntry, _ error) error {
					if !d.IsDir() && md.IsMarkdownFile(path) {
						return md.ProcessMDFile(path, o.Write)
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

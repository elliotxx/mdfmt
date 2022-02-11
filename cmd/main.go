package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/88250/lute"
	"github.com/elliotxx/mdfmt/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cmdShort   = i18n.T(`A Markdown formatter based on CommonMark`)
	cmdLong    = templates.LongDesc(i18n.T(`Like gofmt, but for Markdown. Based on CommonMark.`))
	cmdExample = templates.Examples(i18n.T(`
		# Show version info
		mdfmt -w README.md
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
		Use:          "mdfmt",
		Short:        cmdShort,
		Long:         cmdLong,
		Example:      cmdExample,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Show version info
			if o.ShowVersion {
				fmt.Println(version.ReleaseVersion())
				return nil
			}

			if len(args) == 0 {
				// if err := processFile("<standard input>", os.Stdin, os.Stdout, true); err != nil {
				// 	report(err)
				// }
				return nil
			}

			content, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
			formatContent := luteEngine.FormatStr("md", string(content))

			if o.Write {
				err := ioutil.WriteFile(args[0], []byte(formatContent), 0o644)
				if err != nil {
					return err
				}
			} else {
				fmt.Println(formatContent)
			}

			return nil
		},
	}

	rootCmd.Flags().BoolVarP(&o.ShowVersion, "version", "V", false, "Show version info")
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
}

package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/gh-license/internal/git"
	"github.com/spenserblack/gh-license/internal/license"
	"github.com/spf13/cobra"
)

var (
	// rootOutput is the output path for the root command.
	rootOutput string
	// rootStdout tells the root command to write to stdout.
	rootStdout bool
)

var rootCmd = &cobra.Command{
	Use:   "gh-license <SPDX-ID>",
	Short: "Fetch and write licenses",
	Long: heredoc.Doc(`
		Fetch and write a license file.

		<SPDX-ID> should be the lowercase SPDX ID of the license.
	`),
	Example: "gh license mit",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stdout := cmd.OutOrStdout()
		spdxId := args[0]
		git, err := git.Default()
		if err != nil {
			return err
		}
		ctx, err := license.DefaultContext(git)
		if err != nil {
			return err
		}
		text, err := license.Get(spdxId, ctx)
		if err != nil {
			return err
		}

		if rootStdout {
			fmt.Fprintln(stdout, text)
		} else {
			err = os.WriteFile(rootOutput, []byte(text), 0666)
			if err != nil {
				return err
			} else {
				fmt.Fprintf(stdout, "Created %s\n", rootOutput)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&rootOutput, "output", "o", "LICENSE", "Where the license should be written to")
	rootCmd.Flags().BoolVar(&rootStdout, "stdout", false, "Write license text to stdout instead of a file")
}

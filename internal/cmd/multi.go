package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/gh-license/internal/git"
	"github.com/spenserblack/gh-license/internal/license"
	"github.com/spf13/cobra"
)

// multiSuffixes maps SPDX IDs to a suffix to use for the license.
var multiSuffixes map[string]string

var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Add multiple licenses",
	Long: heredoc.Doc(`
		Add multiple licenses to a project.

		This command will automatically append suffixes in the form LICENSE-<SUFFIX>.
		When the suffix is unknown, it will use the SPDX ID in lowercase as the suffix.
	`),
	Example: "gh license multi apache-2.0 mit",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stdout := cmd.OutOrStdout()
		git, err := git.Default()
		if err != nil {
			return err
		}
		ctx, err := license.DefaultContext(git)
		if err != nil {
			return err
		}

		for _, spdxId := range args {
			text, err := license.Get(spdxId, ctx)
			if err != nil {
				return err
			}
			suffix, ok := multiSuffixes[spdxId]
			if !ok {
				suffix = spdxId
			}

			filename := "LICENSE-" + suffix
			err = os.WriteFile(filename, []byte(text), 0666)
			if err != nil {
				return err
			}
			fmt.Fprintf(stdout, "Created %s\n", filename)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(multiCmd)
	multiCmd.PersistentFlags().StringToStringVarP(&multiSuffixes, "suffixes", "s", map[string]string{
		"apache-2.0": "APACHE",
		"mit":        "MIT",
	}, "Define a map of lowercase SPDX IDs to license file suffixes")
}

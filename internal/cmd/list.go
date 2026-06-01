package cmd

import (
	"fmt"

	"github.com/spenserblack/gh-license/internal/license"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists licenses and their SPDX IDs in lowercase",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		stdout := cmd.OutOrStdout()
		licenses, err := license.List()
		if err != nil {
			return err
		}
		// TODO Make output more readable
		for key, name := range licenses {
			fmt.Fprintf(stdout, "%s\t%s\n", key, name)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

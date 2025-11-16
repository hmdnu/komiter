/*
Copyright © 2025 hmdnu
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	IsCopyToClipboard bool
)

var rootCmd = &cobra.Command{
	Use:   "committer",
	Short: "Use AI to write your git commits",
	Long:  `Use AI to write your git commits`,
	Run:   RunCommiter,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&IsCopyToClipboard, "clipboard", "c", false, "Copy to clipboard")
}

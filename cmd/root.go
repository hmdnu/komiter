package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	IsCopyToClipboard bool
	CommitScope       string
)

var rootCmd = &cobra.Command{
	Use:   "komiter",
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
	rootCmd.Flags().StringVarP(&CommitScope, "commit scope", "s", "", "Provide commit scope")
}

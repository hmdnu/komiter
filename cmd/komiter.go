package cmd

import (
	"fmt"

	"github.com/hmdnu/komiter/internal/ai"
	"github.com/hmdnu/komiter/internal/clip"
	"github.com/hmdnu/komiter/internal/git"
	"github.com/hmdnu/komiter/pkg/prompt"
	"github.com/spf13/cobra"
)

func RunCommiter(command *cobra.Command, args []string) {
	gitDiff, err := git.GetGitDiff()
	if err != nil {
		fmt.Println("Error: failed reading git diff", err)
		return
	}
	prompt := prompt.GenerateCommitPrompt(gitDiff)
	response, err := ai.HttpRequest(prompt)
	if err != nil {
		fmt.Println("Error: failed generating response", err)
		return
	}
	if IsCopyToClipboard {
		clip.CopyToClipboard(response)
		fmt.Println(response)
		fmt.Println("\nCommit message copied to clipboard")
		return
	}
	fmt.Println(response)
}

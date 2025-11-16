package cmd

import (
	"fmt"

	"github.com/hmdnu/committer/internal/ai"
	"github.com/hmdnu/committer/internal/clip"
	"github.com/hmdnu/committer/internal/git"
	"github.com/hmdnu/committer/pkg/prompt"
	"github.com/spf13/cobra"
)

func RunCommiter(command *cobra.Command, args []string) {
	gitDiff, err := git.GetGitDiff()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println("Commit message copied to clipboard")
		return
	}
	fmt.Println(response)
}

package git

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func GetGitDiff() (string, error) {
	command := exec.Command("git", "diff", "--staged")
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		status := strings.Contains(err.Error(), "exit status 129")
		if !status {
			return "", err
		}
		return "", errors.New("Error: not a git repository")
	}

	if stdout.String() == "" {
		return "", errors.New("\nNo files staged")
	}
	return stdout.String(), nil
}

package clip

import (
	"log"

	"github.com/atotto/clipboard"
)

func CopyToClipboard(commitMessage string) {
	if err := clipboard.WriteAll(commitMessage); err != nil {
		log.Fatalf("Error: Copy failed: %v\n", err)
	}
}

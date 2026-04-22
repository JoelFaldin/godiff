package runner

import (
	"fmt"
	"os/exec"
)

func Gitdiff(location string) ([]byte, error) {
	cmd := exec.Command("git", "diff", location)
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("git diff failed: %w", err)
	}

	if len(output) == 0 {
		return nil, fmt.Errorf("no diff output - file may be untracked or unchanged")
	}

	return output, nil
}

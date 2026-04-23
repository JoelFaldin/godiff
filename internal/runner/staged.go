package runner

import (
	"fmt"
	"os/exec"
)

func GitDiffStaged(location string) ([]byte, error) {
	cmd := exec.Command("git", "diff", "--staged", location)
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("git diff --staged failed: %w", err)
	}

	return output, nil
}

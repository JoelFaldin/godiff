package runner

import (
	"fmt"
	"os/exec"
)

func Gitdiff(location string) (rawDiff []byte) {
	cmd := exec.Command("git", "diff", location)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	return output
}

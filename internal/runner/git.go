package runner

import (
	"fmt"
	"os/exec"
)

func Gitdiff(location string) {
	cmd := exec.Command("git", "diff", location)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}

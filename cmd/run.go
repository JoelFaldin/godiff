/*
Copyright © 2026 Joel Faldin joelfaldin@gmail.com
*/
package cmd

import (
	"fmt"
	"godiff/internal/errors"
	"godiff/internal/parser"
	"godiff/internal/renderer"
	"godiff/internal/runner"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the basic command of godiff",
	Long:  `Run and format git diff on the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag, err := cmd.Flags().GetBool("staged")

		if err != nil {
			errors.Print("error processing flag", err.Error())
			return
		}

		location := ""
		if len(args) != 0 {
			location = args[0]
		} else {
			location = "."
		}

		var diff []byte
		var error error

		fmt.Println(flag)
		if flag {
			diff, error = runner.GitDiffStaged(location)
		} else {
			diff, error = runner.Gitdiff(location)
		}

		if error != nil {
			errors.Print("error while running git diff", error.Error())
			return
		}

		parsed, insertions, deletions := parser.Parser(diff)
		renderer.Render(parsed, insertions, deletions)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Local flags (only apply to godiff run):
	runCmd.Flags().Bool("staged", true, "Show differences between files already added to staging area (via git add)")
}

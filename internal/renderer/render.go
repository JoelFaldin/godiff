package renderer

import (
	"fmt"
	"godiff/internal/parser"

	"github.com/charmbracelet/lipgloss"
)

func Render(res []parser.FileDiff) {
	// f, _ := json.MarshalIndent(res, "", " ")
	// result := string(f)
	// fmt.Println(result)

	fileHeader := lipgloss.NewStyle().
		Bold(true).
		Underline(true)

	hunkHeader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("33"))

	addedLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color("42"))

	deletedLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color("196"))

	contextLine := lipgloss.NewStyle().
		Faint(true)

	for i := range res {
		header := fileHeader.Render(res[i].NewPath)
		fmt.Println(header)

		for r := range res[i].Hunks {
			hunk := res[i].Hunks[r]

			hunkHeader1 := hunkHeader.Render("@@ ")
			oldStartHeader := hunkHeader.Render(hunk.OldStart)
			oldCountHeader := hunkHeader.Render(hunk.OldCount)
			newStartHeader := hunkHeader.Render(hunk.NewStart)
			newCountHeader := hunkHeader.Render(hunk.NewCount)
			hunkHeader2 := hunkHeader.Render("@@ ")

			fmt.Println(hunkHeader1, oldStartHeader, oldCountHeader, newStartHeader, newCountHeader, hunkHeader2)

			for l := range hunk.Lines {
				line := hunk.Lines[l]

				if line.Type == 0 {
					reader := contextLine.Render(line.Content)
					fmt.Println(reader)
				} else if line.Type == 1 {
					reader := deletedLine.Render(line.Content)
					fmt.Println(reader)
				} else {
					reader := addedLine.Render(line.Content)
					fmt.Println(reader)
				}
			}
		}
	}
}

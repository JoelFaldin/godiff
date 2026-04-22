package errors

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
var hintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Faint(true)

func Print(msg string, hint string) {
	fmt.Println(errorStyle.Render("x " + msg))

	if hint != "" {
		fmt.Println(hintStyle.Render("→ " + hint))
	}
}

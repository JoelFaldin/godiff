package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func Parser(rawDiff []byte) {
	reader := bytes.NewReader(rawDiff)
	scanner := bufio.NewScanner(reader)

	fmt.Println(string(rawDiff))

	var diffs []FileDiff
	var currentFile *FileDiff
	var currentHunk *Hunk

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "diff --git"):
			if currentFile != nil {
				diffs = append(diffs, *currentFile)
			}
			currentFile = &FileDiff{}

		case strings.HasPrefix(line, "---"):
			if currentFile != nil {
				currentFile.OldPath = strings.TrimPrefix(line, "--- a/")
			}

		case strings.HasPrefix(line, "+++"):
			if currentFile != nil {
				currentFile.NewPath = strings.TrimPrefix(line, "+++ b/")
			}

		case strings.HasPrefix(line, "@@"):
			if currentFile != nil {
				if currentHunk != nil {
					currentFile.Hunks = append(currentFile.Hunks, *currentHunk)
				}
				currentHunk = &Hunk{}

				// Processing old start lines:
				s := line[4:]
				oldStart, _, _ := strings.Cut(s, ",")
				v := line[7:]
				oldCount, _, _ := strings.Cut(v, " ")

				// Processing newstart lines:
				_, n, _ := strings.Cut(line, "+")
				newStart, _, _ := strings.Cut(n, ",")
				b := n[3:]
				newCount, _, _ := strings.Cut(b, " ")

				// currentHunk.OldStart, err = strconv.Atoi(oldStart)
				// currentHunk.OldCount, err2 = strconv.Atoi(oldCount)
				// currentHunk.NewStart, err3 = strconv.Atoi(newStart)
				// currentHunk.NewCount, err4 = strconv.Atoi(newCount)
				fmt.Println(oldStart)
				fmt.Println(oldCount)
				fmt.Println(newStart)
				fmt.Println(newCount)
				// oldCount := 0
				// newStart := 0
				// newCount := 0
			}

		case strings.HasPrefix(line, "+"):
			if currentHunk != nil {
				currentHunk.Lines = append(currentHunk.Lines, Line{
					Type:    Added,
					Content: strings.TrimPrefix(line, "+"),
				})
			}

		case strings.HasPrefix(line, "-"):
			if currentHunk != nil {
				currentHunk.Lines = append(currentHunk.Lines, Line{
					Type:    Removed,
					Content: strings.TrimPrefix(line, "-"),
				})
			}

		default:
			if currentHunk != nil {
				currentHunk.Lines = append(currentHunk.Lines, Line{
					Type:    Context,
					Content: line,
				})
			}
		}
	}

	if currentHunk != nil && currentFile != nil {
		currentFile.Hunks = append(currentFile.Hunks, *currentHunk)
	}

	if currentFile != nil {
		diffs = append(diffs, *currentFile)
	}

	f, _ := json.MarshalIndent(diffs, "", " ")
	fmt.Println(string(f))
}

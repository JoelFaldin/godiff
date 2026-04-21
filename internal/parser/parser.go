package parser

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type parser struct {
	err error
}

func (p *parser) atoi(s string) int {
	if p.err != nil {
		return 0
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		p.err = err
	}

	return val
}

func Parser(rawDiff []byte) []FileDiff {
	reader := bytes.NewReader(rawDiff)
	scanner := bufio.NewScanner(reader)

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

				p := &parser{}
				currentHunk.OldStart = p.atoi(oldStart)
				currentHunk.OldCount = p.atoi(oldCount)
				currentHunk.NewStart = p.atoi(newStart)
				currentHunk.NewCount = p.atoi(newCount)
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

	if currentFile.OldPath == currentFile.NewPath {
		currentFile.IsNew = false
	}

	return diffs
}

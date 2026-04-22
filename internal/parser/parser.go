package parser

import (
	"bufio"
	"bytes"
	"strings"
)

func finalizeFile(f *FileDiff) {
	if f.NewPath == "/dev/null" {
		f.IsDeleted = true
	}

	if f.OldPath == "/dev/null" {
		f.IsNew = true
	}
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
				if currentHunk != nil {
					currentFile.Hunks = append(currentFile.Hunks, *currentHunk)
				}
				finalizeFile(currentFile)
				diffs = append(diffs, *currentFile)
				currentHunk = nil
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

				fields := strings.Fields(line)

				h := fields[1]
				res1 := strings.Split(h, ",")
				currentHunk.OldStart = strings.TrimPrefix(res1[0], "-")
				currentHunk.OldCount = res1[1]

				s := fields[2]
				res2 := strings.Split(s, ",")
				currentHunk.NewStart = strings.TrimPrefix(res2[0], "+")
				currentHunk.NewCount = res2[1]
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

	if currentFile != nil {
		if currentHunk != nil {
			currentFile.Hunks = append(currentFile.Hunks, *currentHunk)
		}
		finalizeFile(currentFile)
		diffs = append(diffs, *currentFile)
	}

	return diffs
}

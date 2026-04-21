package parser

type LineType int

const (
	Context LineType = iota
	Removed
	Added
)

type Line struct {
	Type    LineType
	Content string
}

type Hunk struct {
	OldStart string
	OldCount string
	NewStart string
	NewCount string
	Lines    []Line
}

type FileDiff struct {
	OldPath   string
	NewPath   string
	Hunks     []Hunk
	IsNew     bool
	IsDeleted bool
	IsRenamed bool
}

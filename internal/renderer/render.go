package renderer

import (
	"encoding/json"
	"fmt"
	"godiff/internal/parser"
)

func Render(res []parser.FileDiff) {
	fmt.Println("Hello from render :D")

	f, _ := json.MarshalIndent(res, "", " ")
	result := string(f)
	fmt.Println(result)
}

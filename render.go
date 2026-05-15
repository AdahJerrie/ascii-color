package main

import (
	"strings"
)

func RenderLine(input, substring, color string, banner map[rune][]string) []string {
	output := []string{}

	start := strings.Index(input, substring)
	full := start + len(substring)
	if start == -1 {
		// substring not found, color nothing
		start, full = 0, 0
	} else if substring == "" {
		// no substring specified, color everything
		start, full = 0, len(input)
	}
	//fmt.Printf("DEBUG: input=%q substring=%q color=%q start=%d full=%d\n", input, substring, color, start, full)

	for rows := 0; rows < 8; rows++ {
		var result strings.Builder
		for i, char := range input {
			renderline := banner[char][rows]
			if i >= start && i < full {
				result.WriteString(Colorize(renderline, color))
			} else {
				result.WriteString(renderline)
			}
		}
		output = append(output, result.String())
	}
	return output
}

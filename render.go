package main

import (
	"strings"
)

func RenderLine(input, substring, color string, banner map[rune][]string) []string {
	output := []string{}

	type Range struct {
		start, full int
	}

	var ranges []Range
	start := 0

	if substring == "" { //no substring color everything
		ranges = []Range{{0, len(input)}}
	} else {
		for {
			idx := strings.Index(input[start:], substring)
			if idx == -1 {
				break
			}
			allthestart := start + idx
			full := allthestart + len(substring)
			ranges = append(ranges, Range{allthestart, full})
			start = allthestart + len(substring)
		}
	}

	if len(ranges) == 0 { //substring not found color nothing
		ranges = []Range{{0, 0}}
	}
	//fmt.Printf("DEBUG: input=%q substring=%q color=%q start=%d full=%d\n", input, substring, color, start, full)

	inRange := func(i int) bool {
		for _, r := range ranges {
			if i >= r.start && i < r.full {
				return true
			}
		}
		return false
	}

	for rows := 0; rows < 8; rows++ {
		var result strings.Builder
		for i, char := range input {
			renderline := banner[char][rows]
			if inRange(i) {
				result.WriteString(Colorize(renderline, color))
			} else {
				result.WriteString(renderline)
			}
		}
		output = append(output, result.String())
	}
	return output
}

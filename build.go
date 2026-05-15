package main

import "strings"

func BuildArt(input, sub, colour string, banner map[rune][]string) string {
	if input == "" {
		return input
	}

	split := SplitInput(input)
	var result strings.Builder

	for i, lines := range split {
		if lines == "" {
			if i < len(split)-1 {
				continue
			}
			result.WriteString("\n")
		}
		rendline := RenderLine(lines, sub, colour, banner)
		for _, line := range rendline {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
	return result.String()
}

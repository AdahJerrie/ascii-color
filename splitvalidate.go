package main

import (
	"fmt"
	"strings"
)

func SplitInput(input string) []string {
	split := strings.Split(input, "\\n")
	return split
}

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("%c invalid ascii char", char)
		}
	}
	return 0, nil
}

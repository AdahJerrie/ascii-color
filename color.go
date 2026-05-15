package main

import "fmt"

const reset = "\033[0m"

func Colorize(text, color string) string {
	colors := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"cyan":    "\033[36m",
		"magenta": "\033[35m",
	}

	colorcode, ok := colors[color]
	if !ok {
		return text
	}

	applied := fmt.Sprintf("%s%s%s", colorcode, text, reset)

	return applied
}

package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(" error reading file")
	}

	if len(file) == 0 {
		return nil, errors.New("inalid file length")
	}

	splitfile := strings.Split(string(file), "\n")

	mapbanner := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9
		if start+9 > len(splitfile) {
			return nil, errors.New("invalid banner length")
		}
		completeline := splitfile[start+1 : start+9]
		mapbanner[char] = completeline
	}
	return mapbanner, nil
}

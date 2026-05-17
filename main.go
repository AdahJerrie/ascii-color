package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	color := ""
	substring := ""
	input := ""
	bannerfile := "./banners/standard.txt"
	var plainargs []string

	if len(os.Args) < 2 {
		fmt.Println("invalid argument length")
		return
	}

	for _, args := range os.Args[1:] {
		if strings.HasPrefix(args, "--color=") {
			color = strings.Split(args, "=")[1]
		} else if args == "standard" || args == "shadow" || args == "thinkertoy" {
			bannerfile = fmt.Sprintf("./banners/%s.txt", args)
		} else {
			plainargs = append(plainargs, args)
		}
	}

	if color == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> something ")
		return
	}

	if len(plainargs) == 1 {
		input = plainargs[0]
	} else if len(plainargs) == 2 {
		input = plainargs[1]
		substring = plainargs[0]
	} else {
		fmt.Println("enter valid input")
	}

	banners, err := LoadBanner(bannerfile)
	if err != nil {
		fmt.Println("error loading bannerfile")
		return
	}

	output := BuildArt(input, substring, color, banners)

	fmt.Println(output)

}

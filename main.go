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
	var plainargs []string

	if len(os.Args) < 2 {
		fmt.Println("invalid argument length")
		return
	}

	for _, args := range os.Args[1:] {
		if strings.HasPrefix(args, "--color=") {
			color = strings.Split(args, "=")[1]
		} else {
			plainargs = append(plainargs, args)
		}
	}

	if len(plainargs) == 1 {
		input = plainargs[0]
	} else if len(plainargs) == 2 {
		input = plainargs[1]
		substring = plainargs[0]
	} else {
		fmt.Println("enter valid input")
	}

	banners := "standard.txt"

	bannerfile, err := LoadBanner(banners)
	if err != nil {
		fmt.Println("error loading bannerfile")
		return
	}

	output := BuildArt(input, substring, color, bannerfile)

	fmt.Println(output)

}

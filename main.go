package main

import (
	"art/cmd"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	STRING, BANNER := args[1], ""
	if len(args) == 3 {
		font := args[2]
		BANNER = cmd.HandleFont(font)
	}
	if BANNER == "" {
		BANNER = "banners/standard.txt"
	}
	art := cmd.GenerateAsciiArt(BANNER, STRING)
	fmt.Print(art)
}

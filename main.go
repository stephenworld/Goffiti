package main

import (
	"goffiti/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Expected 2 arguments")
		return
	}

	input := os.Args[1]

	output := utils.AsciiArt(input)
	fmt.Print(output)
}

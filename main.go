package main

import (
	"fmt"
	"goffiti/utils"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Expected atleast 2 arguments")
		return
	}

	input := os.Args[1]

	output := utils.AsciiArt(input)
	fmt.Print(output)
}

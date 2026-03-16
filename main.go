package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Expected 2 arguments")
		return
	}

	input := os.Args[1]

	output := AsciiArt(input)
	fmt.Print(output)
}

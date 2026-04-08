package main

import (
	"flag"
	"fmt"
	"goffiti/cmd"
	"os"
)

func main() {

	// Optional Flags
	align := flag.String("align", "", "Align STRING")
	color := flag.String("color", "", "Apply color to STRING")
	output := flag.String("output", "", "Add content to the output file")

	// Parse Flag contents
	flag.Parse()

	// Stores arguments that are not flags
	arguments := flag.Args()

	art := ""

	if len(arguments) == 1 {
		art = cmd.GenerateAsciiArt(arguments[0])
		fmt.Print(art)
	}

	if len(arguments) == 2 && len(os.Args) == 3 {
		art := cmd.GenerateAsciiArtFS(arguments[0], arguments[1])
		fmt.Println("Goffiti FS Output Start")
		fmt.Print(art)
		fmt.Println("Goffiti FS Output End")
	}

	if *align != "" {
		if len(arguments) == 2 {
			FSArt := cmd.GenerateAsciiArtFS(arguments[0], arguments[1])
			art = cmd.GenerateAsciiArtJustify(FSArt, *align, cmd.GetTerminalWidth())
			fmt.Println(art)
		}
		art = cmd.GenerateAsciiArtJustify(art, *align, cmd.GetTerminalWidth())
		fmt.Println(art)
	}

	if *color != "" {
	}

	if *output != "" {
		if len(arguments) == 2 {
			art := cmd.GenerateAsciiArtFS(arguments[0], arguments[1])
			os.WriteFile(*output, []byte(art), 0644)
		}
		art = cmd.GenerateAsciiArt(arguments[0])
		os.WriteFile(*output, []byte(art), 0644)
	}
}

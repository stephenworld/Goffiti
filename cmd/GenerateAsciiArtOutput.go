package cmd

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAsciiArtOutput(STR, OUTPUT string) {
	font, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fontLines := strings.Split(string(font), "\n")
	strLines := strings.Split(STR, "\\n")

	var ART strings.Builder

	for _, word := range strLines {
		if word == "" {
			ART.WriteString("\n")
			continue
		}
		for idx := range 8 {
			for _, char := range word {
				charLine := GetCharLine(char, fontLines)
				ART.WriteString(charLine[idx])
			}
			ART.WriteString("\n")
		}
	}

	os.WriteFile(OUTPUT, []byte(ART.String()), 0644)
}

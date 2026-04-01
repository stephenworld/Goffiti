package cmd

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAsciiArtFS(STR, BANNER string) string {
	fontPath := HandleFont(BANNER)
	if fontPath == "" {
		return "Font Doest Exist\n"
	}

	font, err := os.ReadFile(fontPath)
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

	return ART.String()
}

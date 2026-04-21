package utils

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAsciiArt(STRING, BANNER string) string {
	font, err := os.ReadFile(BANNER)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	lines := strings.Split(string(font), "\n")
	strLines := strings.Split(strings.ReplaceAll(STRING, "\\n", "\n"), "\n")

	var art strings.Builder

	for _, word := range strLines {
		if word == "" {
			art.WriteString("\n")
			continue
		}
		for idx := range 8 {
			for _, char := range word {
				if char == '\n' {
					continue
				}
				if char < 32 || char > 126 {
					continue
				}
				charLine := GetCharLine(char, lines)
				art.WriteString(charLine[idx])
			}
			art.WriteString("\n")
		}
	}

	return art.String()
}

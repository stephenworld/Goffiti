package cmd

import (
	"os"
	"strings"
)

func GenerateAsciiArt(BANNER, STRING string) string {
	font, err := os.ReadFile(BANNER)
	if err != nil {
		return err.Error()
	}

	fontLines := strings.Split(string(font), "\n")
	stringLines := strings.Split(STRING, "\\n")

	var art strings.Builder

	for _, line := range stringLines {
		if line == "" {
			art.WriteString("\n")
			continue
		}

		for idx := range 8 {
			for _, char := range line {
				if char < 32 && char > 126 {
					continue
				}
				charLine := GetCharLine(char, fontLines)
				art.WriteString(charLine[idx])
			}

			art.WriteString("\n")
		}
	}

	return art.String()
}

package utils

import (
	"os"
	"strings"
)

func ProcessAscii(TEXT, BANNER string) string {
	font, err := os.ReadFile(BANNER)
	if err != nil {
		return err.Error()
	}

	fontLines := strings.Split(string(font), "\n")
	textLines := strings.Split(strings.ReplaceAll(TEXT, "\\n", "\n"), "\n")

	var art strings.Builder

	for _, word := range textLines {

		if word == "" {
			art.WriteRune('\n')
			continue
		}
		for idx := range 8 {
			for _, char := range word {
				if char < 32 && char > 126 {
					continue
				}
				charLine := GetCharLines(char, fontLines)
				art.WriteString(charLine[idx])
			}
			art.WriteRune('\n')
		}
	}

	return art.String()
}

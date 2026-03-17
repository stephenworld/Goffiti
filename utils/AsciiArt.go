package utils

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(str string) string {
	standardAscii, err := os.ReadFile("utils/standard.ascii")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}

	content := strings.ReplaceAll(string(standardAscii), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	if len(str) <= 0 {
		return ""
	}

	result := ""
	for row := 0; row < 8; row++ {
		for _, char := range str {
			charLines := GetCharLines(char, lines)
			result += charLines[row]
		}
		result += "\n"
	}

	return result
}

func GetCharLines(char rune, lines []string) []string {
	startLine := ((char - 32) * 9) + 1
	return lines[startLine : startLine+8]
}

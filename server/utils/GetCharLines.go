package utils

func GetCharLine(char rune, lines []string) []string {
	startLine := (char-' ')*9 + 1
	return lines[startLine : startLine+8]
}

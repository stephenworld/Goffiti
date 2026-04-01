package cmd

func GetCharLine(CHAR rune, LINES []string) []string {
	start := ((CHAR - ' ') * 9) + 1
	end := start + 8
	return LINES[start:end]
}

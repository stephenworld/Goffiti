package cmd

func GetCharLine(CHAR rune, LINES []string) []string {
	start := ((CHAR - ' ') * 9) + 1
	return LINES[start : start+8]
}

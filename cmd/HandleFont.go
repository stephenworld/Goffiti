package cmd

import "strings"

func HandleFont(BANNER string) string {
	lower := strings.ToLower(BANNER)

	switch lower {
	case "shadow", "standard", "thinkertoy":
		return "banners/" + lower + ".txt"
	default:
		return ""
	}
}

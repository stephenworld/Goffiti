package utils

import "strings"

func HandleFont(s string) string {
	lower := strings.ToLower(s)
	switch lower {
	case "standard", "shadow", "thinkertoy":
		return "server/banners/" + lower + ".txt"
	default:
		return ""
	}
}

package utils

import (
	"strings"
)

func HandleBanner(font string) string {
	low := strings.ToLower(font)
	switch low {
	case "standard", "shadow", "thinkertoy":
		return "server/banner/" + low + ".txt"
	default:
		return ""
	}
}

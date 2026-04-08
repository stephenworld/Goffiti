package cmd

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	// "tput cols" returns "cols"
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin // Connect to current terminal
	out, err := cmd.Output()
	if err != nil {
		return 80 // Fallback
	}

	trimSpace := strings.TrimSpace(string(out))

	width, err := strconv.Atoi(trimSpace)
	if err != nil {
		return 80 //Falback
	}

	return width
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(str string) string {
	standardAscii, err := os.ReadFile("standard.ascii")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}

	content := strings.ReplaceAll(string(standardAscii), "\r\n", "\n")
	if content == "" {
		fmt.Println("Error: File content empty")
		return ""
	}

	result := ""
	/*
		TODO:
		1. Handle each char in str and convert them to ascii art
		2. Each ASCII Arts shouldn't exceed 8 lines
	*/

	return result
}

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetNoteData() (string, string) {
	title := GetUserInput("Title: ")
	content := GetUserInput("Content: ")
	return title, content
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

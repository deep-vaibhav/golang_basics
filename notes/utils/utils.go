package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	SaveToFile() error
}

type Outputtable interface {
	Saver
	Display()
}

func OutputData(data Outputtable) error {
	data.Display()
	return data.SaveToFile()
}

func SaveData(data Saver) error {
	err := data.SaveToFile()

	if err != nil {
		fmt.Println("Save failed.")
		return err
	}

	return nil
}

func GetTodoData() string {
	return GetUserInput("Text: ")
}

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

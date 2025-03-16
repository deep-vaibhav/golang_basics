package main

import (
	"fmt"

	"example.com/notes/note"
	"example.com/notes/utils"
)

func main() {
	title, content := utils.GetNoteData()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()
	err = newNote.SaveToFile()

	if err != nil {
		fmt.Println("Save failed")
		return
	}
}

package main

import (
	"fmt"

	"example.com/notes/note"
	"example.com/notes/todo"
	"example.com/notes/utils"
)

func main() {
	title, content := utils.GetNoteData()
	todoText := utils.GetTodoData()

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	utils.OutputData(todo)
	utils.OutputData(newNote)
}

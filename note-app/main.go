package main

import (
	"fmt"

	"example.com/note/helpers"
	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := note.GetNoteData()
	text := todo.GetTodoData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Display()

	err = helpers.SaveData(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Display()

	err = helpers.SaveData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

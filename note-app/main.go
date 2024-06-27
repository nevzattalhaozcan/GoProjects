package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := note.GetNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Println("saving the note failed")
		return
	}
	fmt.Println("the note has been saved successfully")
}

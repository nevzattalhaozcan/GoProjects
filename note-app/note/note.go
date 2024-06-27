package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/note/helpers"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("%v\n\n%v\n\n", note.Title, note.Content)
}

func GetNoteData() (string, string) {
	title := helpers.GetUserInput("Title:")
	content := helpers.GetUserInput("Content:")
	return title, content
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

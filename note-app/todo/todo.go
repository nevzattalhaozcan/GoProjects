package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/note/helpers"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("title and content cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("%v\n\n", todo.Text)
}

func GetTodoData() string {
	text := helpers.GetUserInput("Text:")
	return text
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

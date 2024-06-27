package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

type Outputtable interface {
	Saver
	Display()
}

func GetUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func SaveData(data Saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	fmt.Println("Saving was successfull")
	return nil
}

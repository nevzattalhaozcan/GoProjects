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

func OutputData(data Outputtable) error {
	data.Display()
	return SaveData(data)
}

// BELOW IS NOT MAIN FEATURE, JUST PRACTICE

func PrintSomething(value any) {
	switch value.(type) { // will check the type of value
	case int:
		fmt.Println("Integer:")
	case float64:
		fmt.Println("Float:")
	case string:
		fmt.Println("String:")
	}
}

func PrintSomethingElse(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Integer:", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("Integer:", stringVal)
		return
	}
}

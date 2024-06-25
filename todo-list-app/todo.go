package main

import (
	"fmt"
	"io"
	"os"
)

const itemListFile = "todolist.txt"

func main() {

	fmt.Println("Welcome to Go To-do List App")

	for {

		fmt.Println("Please choose an option below:")
		fmt.Println("1. Show the list")
		fmt.Println("2. Add new item")
		fmt.Println("3. Delete an item")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayList()
		case 2:
			item := getUserInput()
			writeItemToFile(item)
			fmt.Println("Item successfully addedd!")
		case 3:
			fmt.Println("This option will be available soon...")
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go To-do List App!")
			return
		default:
			fmt.Println("invalid input")
		}

	}
}

func writeItemToFile(text string) {
	formattedText := fmt.Sprintf("%v", text)
	os.WriteFile(itemListFile, []byte(formattedText), 0644)
}

func getTodoList() string {
	file, _ := os.OpenFile(itemListFile, os.O_RDONLY, 0644)
	defer file.Close()
	data, _ := io.ReadAll(file)
	todoList := string(data)
	return todoList
}

func displayList() {
	fmt.Println("--------------------")
	fmt.Println("TODO LIST")
	fmt.Printf("-> %v\n", getTodoList())
	fmt.Println("--------------------")
}

func getUserInput() string {
	var item string
	fmt.Print("Please enter your item: ")
	fmt.Scan(&item)
	return item
}

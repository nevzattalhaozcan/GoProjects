package main

import (
	"fmt"

	"example.com/structs/user"
)

type str string // we can create our custom types using built-in types

func (str) someMethod() {
	// we can use built in types as "receiver" so we can use custom built-in types
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	// appUser := user{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthDate: birthdate,
	// 	createdAt: time.Now(),
	// }

	admin := user.NewAdmin("example@test.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	var somethingCool str = "Cool!"
	fmt.Println(somethingCool)
	somethingCool.someMethod()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

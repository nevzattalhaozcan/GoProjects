package main

import "fmt"

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("some")
}

func printSomething(value interface{}) {
	fmt.Println(value)
}
func printSomethingElse(value any) { // exactly the same as above
	fmt.Println(value)
}

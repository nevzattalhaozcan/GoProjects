package main

import "fmt"

type myMap map[string]string

func (m myMap) output() {
	fmt.Println(m)
}

func main() {

	userMap := map[string]string{}
	userMap["johndoe"] = "john123"
	userMap["marryjane"] = "marry123"
	userMap["foo"] = "1234"
	userMap["bar"] = "5678"

	fmt.Println("--- ALL ITEMS ---")
	for key, val := range userMap {
		fmt.Printf("username: %v | password: %v\n", key, val)
	}

	delete(userMap, "foo")
	delete(userMap, "bar")

	fmt.Println("--- AFTER DELETED ITEMS ---")
	for key, val := range userMap {
		fmt.Printf("username: %v | password: %v\n", key, val)
	}

	//predefinedArray := make([]string, 2, 5)

	predefinedMap := make(myMap, 2)

	predefinedMap.output()

}

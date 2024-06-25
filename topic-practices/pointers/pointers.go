package main

import (
	"fmt"
	"time"
)

func main() {

	year := 1996

	//var pointerForBirth *int

	pointerForBirth := &year

	age := calculateAge(pointerForBirth)

	//var agePointer *int
	agePointer := &age

	fmt.Println("My Birth Year:", year)
	fmt.Println("My Age:", age)
	returnToAdultYears(agePointer)
	fmt.Println("My Adult Years:", age)
	fmt.Println("My Age:", age)

}

func calculateAge(year *int) int {
	return time.Now().Year() - *year
}

func returnToAdultYears(age *int) {
	*age = *age - 18
}

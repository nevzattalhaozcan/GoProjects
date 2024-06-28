package main

import "fmt"

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	//	  Output (print) that array in the command line.
	fmt.Println("----- 1 -----")
	myHobbies := [3]string{"coding", "walking", "reading"}
	fmt.Println("My Hobbies:", myHobbies)

	// 2) Also output more data about that array:
	//	- The first element (standalone)
	//	- The second and third element combined as a new list
	fmt.Println("----- 2 -----")
	fmt.Println("My Hobbies:", myHobbies)
	fmt.Println("First:", myHobbies[0])
	secondAndThirdHobby := myHobbies[1:]
	fmt.Println("Slice of Second and Third:", secondAndThirdHobby)

	// 3) Create a slice based on the first element that contains
	//	  the first and second elements.
	//	  Create that slice in two different ways (i.e. create two slices in the end)
	fmt.Println("----- 3 -----")
	fmt.Println("My Hobbies:", myHobbies)
	firstTwoElements := myHobbies[0:2]
	fmt.Println("Slice of First and Second:", firstTwoElements)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//	  and last element of the original array.
	fmt.Println("----- 4 -----")
	fmt.Println("My Hobbies:", myHobbies)
	secondAndLastElements := firstTwoElements[1:3]
	fmt.Println("Slice of Second and Last Elements:", secondAndLastElements)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	fmt.Println("----- 5 -----")
	myCourseGoals := []string{"be able to read Go code", "be able to write some basic Go code for testing"}
	fmt.Println("My Course Goals:", myCourseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	fmt.Println("----- 6 -----")
	myCourseGoals[1] = "be able to use Go for testing"
	myCourseGoals = append(myCourseGoals, "be able to create projects in the future")
	fmt.Println("My Course Goals:", myCourseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//			 dynamic list of products (at least 2 products).
	//			 Then add a third product to the existing list of products.
	fmt.Println("----- 7 -----")
	type product struct {
		title string
		id    int
		price float64
	}
	products := []product{{title: "egg", id: 1, price: 5.5}, {title: "bread", id: 2, price: 7.5}}
	products = append(products, product{title: "water", id: 3, price: 6.5})
	fmt.Println("Products:", products)

	// Unpacking Arrays to merge
	fmt.Println("----- Unpacking Arrays -----")
	firstArray := []int{1, 2, 3}
	fmt.Println(firstArray)
	secondArray := []int{4, 5, 6}
	fmt.Println(secondArray)
	firstArray = append(firstArray, secondArray...) // use three dots
	fmt.Println(firstArray)

}

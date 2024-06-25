package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const seatsFile = "seats.txt"

var fromLocation string
var toLocation string
var date string
var time string
var seatNumber int

func main() {
	appUi()
}

func appUi() {
	fmt.Println("Welcome to Go Reservation!")
	fmt.Print("From Location: ")
	fmt.Scan(&fromLocation)
	fmt.Print("To Location: ")
	fmt.Scan(&toLocation)
	fmt.Print("Date(dd.mm.yyyy): ")
	fmt.Scan(&date)
	fmt.Print("Time(hh.mm): ")
	fmt.Scan(&time)
	fmt.Println("Choose a seat number")
	displaySeats()
	fmt.Print("Your seat: ")
	fmt.Scan(&seatNumber)
	chooseSeat(seatNumber)
	listTimes()
}

func displaySeats() {
	rows := 6
	columns := 4

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			seatNum := i*columns + j + 1
			if j == 2 {
				fmt.Print("   ")
			}
			for _, seatStr := range strings.Split(readSeats(), ", ") {
				seat64, _ := strconv.ParseInt(seatStr, 8, 0)
				seat := int(seat64)
				if seat == seatNumber {
					fmt.Print("| x |")
				} else {
					fmt.Printf("| %2d |", seatNum)
				}
			}
		}
		fmt.Println()
	}
}

func chooseSeat(seatNum int) {
	file, _ := os.OpenFile(seatsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	seatText := fmt.Sprintf("%v, ", seatNum)
	file.WriteString(seatText)
}

func listTimes() {
	fmt.Printf("Listing trips from %v to %v\n", fromLocation, toLocation)
	for i := 1; i <= 24; i++ {
		if i < 10 {
			fmt.Printf("0%v.00\n", i)
		} else {
			fmt.Printf("%v.00\n", i)
		}
	}
}

func readSeats() string {
	seats, _ := os.ReadFile(seatsFile)
	return string(seats)
}

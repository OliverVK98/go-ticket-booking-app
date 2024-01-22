package main

import (
	"fmt"
	"strings"
)

var appName = "Ticket Booking App"

const availableTicketQuantity = 50

var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTicketQuantity := getUserInput()
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicketQuantity > 0

		if !isValidEmail || !isValidTicketNumber {
			fmt.Println("Invalid user input")
			continue
		}

		if userTicketQuantity > remainingTickets {
			fmt.Printf("We only have %v tickets left\n", remainingTickets)
			continue
		}

		remainingTickets -= userTicketQuantity
		bookings = append(bookings, fmt.Sprintf("%v %v", firstName, lastName))

		printFirstNames()

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicketQuantity, email)

		if remainingTickets < 1 {
			fmt.Println("All tickets are sold out!")
			break
		}
	}
}

func greetUsers() {
	greetingMessage := fmt.Sprintf("Welcome to %s! There are total %d tickets available and %d tickets remaining", appName, availableTicketQuantity, remainingTickets)
	fmt.Println(greetingMessage)
}

func printFirstNames() {
	var firstNames []string
	for _, fn := range bookings {
		firstNames = append(firstNames, strings.Fields(fn)[0])
	}
	fmt.Printf("Current attendee list %v \n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicketQuantity uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter desired ticket quantity: ")
	fmt.Scan(&userTicketQuantity)

	return firstName, lastName, email, userTicketQuantity
}

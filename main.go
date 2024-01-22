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

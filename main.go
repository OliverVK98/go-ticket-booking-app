package main

import (
	"fmt"
	"strconv"
	"strings"
)

var appName = "Ticket Booking App"

const availableTicketQuantity = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

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

		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["userTicketQuantity"] = strconv.FormatUint(uint64(userTicketQuantity), 10)

		bookings = append(bookings, userData)

		printFirstNames()

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicketQuantity, email)

		if remainingTickets < 1 {
			fmt.Println("All tickets are sold out!")
			break
		}
	}
}

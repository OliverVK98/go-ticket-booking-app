package main

import (
	"fmt"
	"strings"
)

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

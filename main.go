package main

import "fmt"

func main() {
	appName := "Ticket Booking App"
	const availableTicketQuantity = 50
	var remainingTickets uint = 50

	greetingMessage := fmt.Sprintf("Welcome to %s! There are total %d tickets available and %d tickets remaining", appName, availableTicketQuantity, remainingTickets)
	fmt.Println(greetingMessage)

	var firstName string
	var lastName string
	var email string
	var userTicketQuantity int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter desired ticket quantity: ")
	fmt.Scan(&userTicketQuantity)

}

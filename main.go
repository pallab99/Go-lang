package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings []string
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
	for {
		var firstName string
		var lasttName string
		var email string
		var userTickets int
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lasttName)

		fmt.Println("Enter your email name:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets

		bookings = append(bookings, firstName+" "+lasttName+",")

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lasttName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are our all the bookings : %v\n", bookings)
	}
}

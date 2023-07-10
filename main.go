package main

import (
	"fmt"
	"strings"
)

func main() {

	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference"
	// we can do the above things as conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// var bookings []string
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		// slices is an abstraction of an Array
		// slices are also index-based and have a size, but is resized when needed

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("slice type: %T\n", bookings)
		// fmt.Printf("slice length: %v\n", len(bookings))

		// userName = "Tom"
		// userTickets = 2
		fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a conformation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		// fmt.Println(conferenceName)
		firstNames := []string{}
		for index, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
		}
		fmt.Printf("These are all our bookings: %v\n", bookings)

	}
}

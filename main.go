package main

import "fmt"

func main() {
	const conferenceTickets = 50

	var conferenceName string = "Go Conference"
	var userName string
	var remainingTickets int = 50
	var userTickets int

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickents and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets attend")

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}

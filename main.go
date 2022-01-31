package main

import "fmt"

func main() {
	const conferenceTickets = 50

	var conferenceName string = "Go Conference"
	var firstName string
	var lastName string
	var email string
	var remainingTickets uint = 50
	var userTickets uint

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickents and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets attend")

	fmt.Println("Enter your fisrt name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booked %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
	fmt.Printf(" %v tickets reamining tickets for %v\n", remainingTickets, conferenceName)
}

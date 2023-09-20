package main

import "fmt"

func main() {
	var eventName = "Go Conference"
	const eventTickets = 50
	var remainingTickets = eventTickets
	// fmt.Println("Welcome to " + eventName + ", Glad to see you!")
	// fmt.Println("We have total of ", eventTickets, " tickets with ", remainingTickets, " are still available.")
	fmt.Printf("Welcome to %v, Glad to see you!\n", eventName)
	fmt.Printf("We have total of %v tickets with %v are still available.", eventTickets, remainingTickets)
}

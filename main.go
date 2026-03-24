package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets  = 50
	var remainingTickets uint = 50
	var bookings []string

	// fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the event")

	for {
		var firstName string
		var lastName string
		var emailID string
		var userTickets uint
		//ask the user for their name
		fmt.Println("Enter your name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&emailID)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		// if userTickets > remainingTickets {
		// 	fmt.Printf("Only %v tickets are left. So you can't book %v tickets.\n", remainingTickets, userTickets)
		// 	continue
		// }
		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName,lastName, userTickets, emailID)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) 
				firstNames = append(firstNames, names[0])
			}
			
			fmt.Printf("The First names of all the our bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				//end program
				fmt.Println("Conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("Only %v tickets are left. So you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}

}

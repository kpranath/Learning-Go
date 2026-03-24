package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"
const conferenceTickets  = 50
var remainingTickets uint = 50
var bookings []string

func main() {
	

	// fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers()


	for {
		firstName, lastName, emailID, userTickets := getUserInput()

		
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, emailID, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicket {
			bookTicket(userTickets, firstName, lastName, emailID)
			
			//call function to print first names
			firstNames := getFirstNames()

			fmt.Printf("The First names of all the our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("Invalid name. Pleass Try again")
			} 
			if !isValidEmail{
				fmt.Println("Invalid email, Please try again")
			}
			if !isValidTicket {
				fmt.Println("Invalid ticket number, Please try again")
			}
		}
	}

}

//Greet the user
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the event")
}

//function to print First Names
func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) 
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

//Function to check user input and validate correct type
func validateUserInput(firstName string, lastName string, emailID string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailID, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail,isValidTicket
}

//Function to get user input
func getUserInput() (string, string,string,uint){
	var firstName string
		var lastName string
		var emailID string
		var userTickets uint
	
		fmt.Println("Enter your name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&emailID)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
		return firstName, lastName,emailID,userTickets
}

//Function to book ticket
func bookTicket(userTickets uint, firstName string, lastName string, emailID string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName,lastName, userTickets, emailID)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
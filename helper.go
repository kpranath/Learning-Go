package main

import "strings"


//Function to check user input and validate correct type
func validateUserInput(firstName string, lastName string, emailID string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailID, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail,isValidTicket
}
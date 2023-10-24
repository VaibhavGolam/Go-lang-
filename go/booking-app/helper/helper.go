package helper

import ("strings")

func ValidateUserInput(firstName string, LastName string,userEmail string,userTicket uint,remainingTickets uint) (bool,bool,bool)  {
	isNameValid := len(firstName) >=2 && len(LastName) >= 2
	isEmailValid := strings.Contains(userEmail,"@") 
	isUserTicketValid := userTicket > 0 && userTicket <= remainingTickets
	return isNameValid, isEmailValid, isUserTicketValid 
}

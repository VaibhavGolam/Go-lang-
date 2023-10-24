package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go conference"

var conferenceTickets = 50         //fixed tickets
var remainingTickets uint = 50     //remaining  tickets
var bookings = make([]UserData, 0) //storage of name.. using slice

type UserData struct {
	firstName       string
	LastName        string
	userEmail       string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	//greets user
	greetUser()
	//takes inputs from user
	firstName, LastName, userEmail, userTicket := getUserInput()

	//input Validation
	isNameValid, isEmailValid, isUserTicketValid := helper.ValidateUserInput(firstName, LastName, userEmail, userTicket, remainingTickets)

	if isNameValid && isEmailValid && isUserTicketValid {

		bookTicket(userTicket, firstName, LastName, userEmail)
		wg.Add(1)
		go sendTicket(userTicket, firstName, LastName, userEmail)

		//prints all first names of booking
		firstNames := getFirstNames()
		fmt.Printf("First names of bookings are : %v\n", firstNames)

		var noTicketRemaining bool = remainingTickets == 0
		if noTicketRemaining {
			//program ends
			fmt.Println("All the tickets are sold out, come back later")
			//break
		}
	} else {
		if !isNameValid {
			fmt.Println("Entered firstname or last name is very short.")
		}
		if !isEmailValid {
			fmt.Println("email adress you entered doesnt contain '@' sign ")
		}
		if !isUserTicketValid {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We hav total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var LastName string
	var userEmail string
	var userTicket uint

	fmt.Println("Enter your first and last name:")
	fmt.Scan(&firstName, &LastName)

	fmt.Println("Enter your email :")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets :")
	fmt.Scan(&userTicket)

	return firstName, LastName, userEmail, userTicket
}

func bookTicket(userTicket uint, firstName string, LastName string, userEmail string) {
	remainingTickets -= userTicket

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		LastName:        LastName,
		userEmail:       userEmail,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData) //using map
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket. you will receive a confirmation mail at %v \n", firstName, LastName, userTicket, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTicket uint, firstName string, LastName string, userEmail string) {
	time.Sleep(25 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, LastName)
	fmt.Println("################################################")
	fmt.Printf("sending ticket :\n %v \n to email address %v \n", ticket, userEmail)
	fmt.Println("################################################")
	wg.Done()
}

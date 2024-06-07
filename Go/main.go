package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTicket int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for remainingTickets > 0 {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isvalidemail, isValidTicketCount := validUserInput(firstName, lastName, email, userTicket)

		if isValidName && isvalidemail && isValidTicketCount {

			if remainingTickets == 0 {
				fmt.Printf("Sorry %v, all tickets are sold out come again next time\n", firstName)
				break
			}

			if remainingTickets < userTicket {
				fmt.Printf("Only %v tickets are availabe for %v , sorry you can't book %v tickets\n", remainingTickets, conferenceName, userTicket)
				continue
			}

			remainingTickets = remainingTickets - userTicket

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTicket,
			}

			bookings = append(bookings, userData)

			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)

			firstNames := getFirstName()

			fmt.Printf("Thank you %v %v for booking %v tickets , confirmation email will be sent to you at %v\n", firstName, lastName, userTicket, email)

			fmt.Printf("%v tickets are available now for %v\n", remainingTickets, conferenceName)

			fmt.Printf("these are all our bookings %v\n", firstNames)

			fmt.Print(bookings)
			fmt.Print("\n")

		} else {
			if !isValidName {
				fmt.Println("please enter valid name")
			}
			if !isvalidemail {
				fmt.Println("please enter valid email address")
			}
			if !isValidTicketCount {
				fmt.Println("please enter valid number of ticket you want to book")
			}
		}

	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and still %v tickets are available\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket here")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidemail := strings.Contains(email, "@")
	isValidTicketCount := userTicket > 0
	return isValidName, isvalidemail, isValidTicketCount
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want to book")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v\n", userTicket, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("sending tickets: %v to email address %v\n", tickets, email)
	fmt.Println("#####################\n")
	wg.Done()
}

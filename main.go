package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint8 = 50

var conferenceName = "GO Conference"
var remainingTickets uint8 = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

var wg = sync.WaitGroup{}

func main() {
	defer wg.Wait()

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our", conferenceName, "is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname and Lastname length must be greater than 1")
			}
			if !isValidEmail {
				fmt.Println("Email must contain '@' and '.'")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}
	}
}

func greetUsers() {
	// consoleLog(fmt.Sprintf("Welcome to %v ticket booking application\n", conferenceName), "red", "", true, false)
	consoleLog(ConsoleOptions{
		msg:      fmt.Sprintf("Welcome to %v ticket booking application", conferenceName),
		bgColor:  "bgCyan",
		isBold:   true,
		isBanner: true,
	})
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets are available")
	// consoleLog("Book your tickets here\n", "blue", "", true, false)
	consoleLog(ConsoleOptions{
		msg:    "Book your tickets here",
		color:  "blue",
		isBold: true,
	})
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint8, firstName string, lastName string, email string) {
	defer wg.Done()

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
}

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/raptordev21/console"
	"github.com/raptordev21/console/colors"
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
			console.Info(fmt.Sprintf("The first names of bookings: %v\n", firstNames))

			if remainingTickets == 0 {
				console.Warn(fmt.Sprint("Our ", conferenceName, " is booked out. Come back next year."))
				break
			}
		} else {
			if !isValidName {
				opts := console.LogOptions{Msg: "Firstname and Lastname length must be greater than 1", Color: colors.Color.Red}
				console.Log(opts)
			}
			if !isValidEmail {
				opts := console.LogOptions{Msg: "Email must contain '@' and '.'", Color: colors.Color.Red}
				console.Log(opts)
			}
			if !isValidTicketNumber {
				opts := console.LogOptions{Msg: fmt.Sprintf("We only have %v tickets remaining, so you can't book %v tickets", remainingTickets, userTickets), Color: colors.Color.Red}
				console.Log(opts)
			}
		}
	}
}

func greetUsers() {
	console.Log(console.LogOptions{Msg: fmt.Sprintf("Welcome to %v ticket booking application", conferenceName), BgColor: colors.Color.BgCyan, IsBold: true, IsBanner: true})
	console.Log(console.LogOptions{Msg: fmt.Sprint("We have total of ", conferenceTickets, " tickets and ", remainingTickets, " tickets are available"), Color: colors.Color.Blue, IsBold: true})
	console.Log(console.LogOptions{Msg: "Book your tickets here:", Color: colors.Color.Blue, IsBold: true})
	fmt.Println()
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

	console.Log(console.LogOptions{Msg: "Enter your first name:", Color: colors.Color.Magenta})
	fmt.Scan(&firstName)

	console.Log(console.LogOptions{Msg: "Enter your last name:", Color: colors.Color.Magenta})
	fmt.Scan(&lastName)

	console.Log(console.LogOptions{Msg: "Enter your email:", Color: colors.Color.Magenta})
	fmt.Scan(&email)

	console.Log(console.LogOptions{Msg: "Enter number of tickets:", Color: colors.Color.Magenta})
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
	console.Info(fmt.Sprintf("List of bookings: %v\n", bookings))

	console.Success(fmt.Sprintf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email))
	console.Info(fmt.Sprintf("%v tickets remaining for %v\n", remainingTickets, conferenceName))
}

func sendTicket(userTickets uint8, firstName string, lastName string, email string) {
	defer wg.Done()

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets to %v %v", userTickets, firstName, lastName)
	console.Log(console.LogOptions{Msg: fmt.Sprintf("Sending %v to email address %v", ticket, email), BgColor: colors.Color.BgGreen, IsBold: true, IsBanner: true})
}

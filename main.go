package main

import (
	"fmt"
	"strconv"
	"strings"
)

const movieTickets = 30

var remainingTickets uint = 30
var bookings []string

func main() {
	fmt.Println("\nWelcome to movie ticket booking app!")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", movieTickets, remainingTickets)

	for {
		firstName, lastName, phoneNumber, email, userTickets := userInput()
		correctName, correctNumber, correctEmail, correctTickets := userVerivication(firstName, lastName, phoneNumber, email, userTickets)

		if correctName && correctNumber && correctEmail && correctTickets {
			listOfSoldTickets := buyTicket(userTickets, firstName, lastName, phoneNumber, email)
			buyerNames(listOfSoldTickets)

			if remainingTickets == 0 {
				fmt.Println("We're sorry, tickets are all sold out.")
				break
			}
		} else {
			if !correctName {
				fmt.Println("Firstname and/or lastname entered incorrectly.")
			}
			if !correctEmail {
				fmt.Println("Email address entered incorrectly and does't contain @.")
			}
			if !correctNumber {
				fmt.Println("Phone number you entered incorrectly. The number must contain at least 5 digits.")
			}
			if !correctTickets {
				fmt.Println("Incorrect entry of information in the field for purchasing tickets.")
			}
		}

	}
}

func userInput() (string, string, uint, string, uint) {
	var firstName string
	var lastName string
	var phoneNumber uint
	var email string
	var userTickets uint

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your lastname: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your phonenumber: ")
	fmt.Scan(&phoneNumber)

	fmt.Print("Enter your email adress: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, phoneNumber, email, userTickets
}

func userVerivication(firstName string, lastName string, phoneNumber uint, email string, userTickets uint) (bool, bool, bool, bool) {
	correctName := len(firstName) >= 2 && len(lastName) >= 2
	correctNumber := len(strconv.Itoa(int(phoneNumber))) > 5
	correctEmail := strings.Contains(email, "@")
	correctTickets := userTickets > 0 && userTickets <= remainingTickets
	return correctName, correctNumber, correctEmail, correctTickets
}

func buyTicket(userTickets uint, firstName string, lastName string, phoneNumber uint, email string) []string {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("%v %v, thank you for buying tickets! Wait for confirmation of the purchase of %v tickets to your phone %v and email adress %v", firstName, lastName, userTickets, phoneNumber, email)
	fmt.Printf("\n%v tickets are available for purchase.", remainingTickets)
	return bookings
}
func buyerNames(bookings []string) {
	buyerNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		buyerNames = append(buyerNames, names[0])
	}
	fmt.Printf("\nList of buyers: %v\n", buyerNames)
}

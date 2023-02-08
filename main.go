package main

import (
	"fmt"
	"strconv"
	"strings"
)

const movieTickets = 30

var remainingTickets uint = 30
var bookings = make([]Customer, 0)

type Customer struct {
	firstName  string
	lastName   string
	phoneNumer uint
	email      string
	tickets    uint
}

func main() {
	fmt.Println("\nWelcome to movie ticket booking app!")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", movieTickets, remainingTickets)

	for {
		firstName, lastName, phoneNumber, email, customerTickets := userInput()
		correctName, correctNumber, correctEmail, correctTickets := userVerivication(firstName, lastName, phoneNumber, email, customerTickets)

		if correctName && correctNumber && correctEmail && correctTickets {
			buyTicket(customerTickets, firstName, lastName, phoneNumber, email)
			buyerNames()

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
	var customerTickets uint

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your lastname: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your phonenumber: ")
	fmt.Scan(&phoneNumber)

	fmt.Print("Enter your email adress: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&customerTickets)
	return firstName, lastName, phoneNumber, email, customerTickets
}

func userVerivication(firstName string, lastName string, phoneNumber uint, email string, customerTickets uint) (bool, bool, bool, bool) {
	correctName := len(firstName) >= 2 && len(lastName) >= 2
	correctNumber := len(strconv.Itoa(int(phoneNumber))) > 5
	correctEmail := strings.Contains(email, "@")
	correctTickets := customerTickets > 0 && customerTickets <= remainingTickets
	return correctName, correctNumber, correctEmail, correctTickets
}

func buyTicket(customerTickets uint, firstName string, lastName string, phoneNumber uint, email string) {
	remainingTickets = remainingTickets - customerTickets

	customerData := Customer{
		firstName:  firstName,
		lastName:   lastName,
		phoneNumer: phoneNumber,
		email:      email,
		tickets:    customerTickets,
	}

	bookings = append(bookings, customerData)
	fmt.Printf("Data list: %v", bookings)

	fmt.Printf("\n%v %v, thank you for buying tickets! Wait for confirmation of the purchase of %v tickets to your phone %v and email adress %v", firstName, lastName, customerTickets, phoneNumber, email)
	fmt.Printf("\n%v tickets are available for purchase.", remainingTickets)
}
func buyerNames() {
	buyerNames := []string{}
	for _, booking := range bookings {
		buyerNames = append(buyerNames, booking.firstName)
	}
	fmt.Printf("\nList of buyers: %v\n", buyerNames)
}

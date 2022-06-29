package main 
import (
	"fmt"
	"booking-app/helper"
	"time"
	"sync"
)
// Multiple functions can have access to a Variable once they are in a global scope; They are called Package Level Variables
// Package level variables cannot be Created using := syntax
var conferenceName string = "Go Conference"
//  This is a constant but cannot be changed
const conferenceTickets uint = 50 
// Constants cant be defined with :=, Take note; Does not work too if you want to define type for a variable
var remainingTickets = conferenceTickets
// This creates an Empty List of Maps; Has an Initial dynamic size of 0
var bookings = make([]UserData, 0)

// Structs lets us define key value pairs but with mixed data types
//  Struct gives us a structure by listing all properties it should have
// They can be compared to classes
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}
var wg = sync.WaitGroup{}
// package called main
// we have to define the entry point of our application
// A program can only have one main function
// Now, our application belongs to a package
func main () {
	
	// Arrays in Go have fixed size [Take note]
	// A slice is an abstraction of an Array , More flexible and powerful; a slice does not have the fixed length
	// To add to a slice append(bookings, value); Getting value from slice is same as Array
	// len is used to check size of a list or of a character 
	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName ,isValidEmail , isValidTicketNumber  := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber  {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("First names of bookings are %v\n ", firstNames)
			// Loops are simplified in GoLang; No diff types of loops 
			// Range allows us to iterate over elements for different data structures 
			if  remainingTickets == 0 {
				fmt.Println("Conference booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname or lastname is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is Invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number")
			}
			continue
		}
	}
	wg.Wait()
}
// Everything is organised into packages
// Think of the packages as containers of various functionalities
// to run a go file "go run main.go"

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v and tickets remaining is %v ", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() [] string {
	firstNames := [] string {}
	// An Underscore (_) is a blank Identifier used to ignore a variable we do not want to use
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}



func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// A pointer is a variable that points to the memory address of another variable ; Pointers is used in only C and C++
	fmt.Println("Whats your first name? ")
	fmt.Scan(&firstName)

	fmt.Println("Whats your last name? ")
	fmt.Scan(&lastName)

	fmt.Println("Whats your email? ")
	fmt.Scan(&email)

	fmt.Println("How  many tickets do you want?")
	fmt.Scanln(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// Make is used to create an empty Map
	// We cannot have mixed Data types as Values in Go
	// strconv Package very nice 
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for bookking %v tickets, you would receive a confirmation email on %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("Your remaining tickets is %v\n-----------------------------\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Lets assume it takes 10 seconds for email to send
	time.Sleep(10 * time.Second) 
	fmt.Println("-----------------")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("-----------------")
	wg.Done()
}
// We need concurrency in our Applications, 

// Green Thread in Go is an abstraction of the OS thread, Its cheaper and more lightweight, less memory space
package helper
// Go is organised into Packages
// A package is a collection of Go files 
import "strings"

// The command go run . runs all the packages in the Go folder
// A go application can have multiple packages

// Variables and Functions defined outrside any function can be accessed in all other files within the same package
// To be able to export a function, We should make it start with an Uppercase letter, eg a fmt package functions start with capital letters
// You can also Export Variables by capitalising them
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && remainingTickets >= userTickets
	return isValidName ,isValidEmail , isValidTicketNumber 
}
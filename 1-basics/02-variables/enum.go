package main

import "fmt"

// doesnt exist enum in go, so we need to create a new type and a set of constants that use this type
// create a new type for each enum is good to avoid comparing 2 different enums and the system thinks that's the same thing
type Status uint8

const(
	OPEN Status = iota // the first constant needs to receive the type os all constants. the compiler will set all consts with the same type. iota is a keyword to set int values for all consts starting at 0
	PENDING // is 1 because of iota
	APPROVED // is 2 because of iota
	DENIED // is 3 because of iota
	CANCELED // is 4 because of iota
)

// create a function called automatically to convert the enum to string without we need to do anything.
// The (s Status) before the name add the method inside the type so it can be called automatically when anything tries to convert
func (s Status) String()string{
	switch s {
	case OPEN:
		return "Open"
	case PENDING:
		return "Pending"
	case APPROVED:
		return "Approved"
	case DENIED:
		return "Denied"
	case CANCELED:
		return "Canceled"
	default:
		return "Unknown"
	}
}

func main(){
	fmt.Println(PENDING) // prints 1 without the String method, with the method prints "Pending"
}
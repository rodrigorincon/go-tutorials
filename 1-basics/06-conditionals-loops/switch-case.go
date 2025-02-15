package main

import "fmt"

func someFunc()int{
	fmt.Println("entered in someFunc")
	return 13
}

func main(){
	// possible case values
	number1 := 11
	number2 := 12
	switch number1 {
	case 10:
		fmt.Println("number1 is 10")
	case 11:
		fmt.Println("number1 is 11")
	case number2: // we can use variables in case clause
		fmt.Println("number1 is equals to number2")
	case someFunc(): // this function is not called if switch enters in some of cases before
		fmt.Println("number1 is equals to someFunc method return")
	default:
		fmt.Println("number1 is another value")
	}

	// if also can be written with parentesis. It's optional
	switch(number1){
	case 10:
		fmt.Println("number1 is 10")
	default:
		fmt.Println("number1 is another value")
	}
	
	// Switch with no condition it's the same as switch TRUE. It'll check all cases in the order they are and stop if enter in someone
	// it's a clean way to write a many of if-else if-else if-else.
	number1 = 10
	switch {
	case number1 < 5:
		fmt.Println("Less than 5")
	case number1 < 10:
		fmt.Println("less than 10")
	case number1 < 20:
		fmt.Println("less than 20")
	default:
		fmt.Println("greather than 20")
	}

	// switch-case without case
	switch(number1){
	default:
		fmt.Println("number1 is some value")
	}
	
	// switch-case without default
	switch(number1){
	case 5:
		fmt.Println("number1 is 10")
	}

	// short var declaration
	switch season := "spring"; season { // we can declare a var in the switch verification and use it only in the swith-case block
	  case "spring":
	    fmt.Println("it's spring")
	  case "summer":
	    fmt.Println("it's summer")
	}
}
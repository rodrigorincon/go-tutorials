package main

import "fmt"

func function1(val bool)bool{
	fmt.Println("function 1")
	return val
}
func function2(val bool)bool{
	fmt.Println("function 2")
	return val
}

func main(){
	// comparing statements
	number1 := 15
	number2 := 10

	if number1 == number2 {
		fmt.Println("number1 and number2 are equal")
	} 
	if number1 != number2 {
		fmt.Println("number1 and number2 are different")
	}
	if(number1 > 5 && number2 > 5){
		fmt.Println("both numbers are greater than 5")
	} 
	//if( 30 > number1 > 5 ){} // this type of comparison doesnt work!!
	
	if(number1 > 10 || number2 > 10){
		fmt.Println("one of the numbers are greater than 10")
	}

	// not statement
	boolValue := true
	if(boolValue){
		fmt.Println("it's true")
	}
	boolValue = false
	if(!boolValue){
		fmt.Println("now it's false")
	}

	// testing execution order in && and || statements
	if( function1(true) && function2(true) ){} // prints both messages
	fmt.Println("-----")
	if( function1(false) && function2(true) ){} // prints only function1 msg
	fmt.Println("-----")
	if( function1(false) || function2(false) ){} // prints both messages

	// OR, AND and NOT statements doesnt exist in Go
	// if(number1 > 5 and number2 > 5){}
	// if(number1 > 5 or number2 > 5){}
	// if(not boolValue){}
}
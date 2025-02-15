package main

import "fmt"

func main(){ 
	// GOLANG DOESNT HAVE TERNARY!!!!!

	
	// if doesnt need parentesis
	number1 := 15
	number2 := 10
	if number1 > number2 {
		fmt.Println("number1 is greater than number2")
	} else {
		fmt.Println("number1 is less than number2")
	}

	// if also can be written with parentesis. It's optional
	if(number1 > number2){
		fmt.Println("number1 is greater than number2")
	} else {
		fmt.Println("number1 is less than number2")
	}
	
	// the use of {} is required for if and else clause/ The follow example doenst compile
	// if number1 > number2
	//	 fmt.Println("something")
	
	// single line clause and block works, but the {} is necessary
	if(number1 > number2){ fmt.Println("something") }

	// else-if works
	if number1 > 18{
		fmt.Println("it is legal age")
	} else if number1 >= 13 { // doesnt exist elsif or elif keywords
		fmt.Println("it is teenager")
	} else{
		fmt.Println("it is kid")
	}

	// short var declaration
	if age := 55; age < 50 { // we can create a var in the if verification and it'll exist only inside the if block
		fmt.Println("age var only exist inside the if block:", age)
	}else{
		fmt.Println("age var only exist inside the if-else block:", age)
	}

	// we can't use only a not boolean var alone in a if statement
	// if(number1){} // doesnt work!
	// if(!number1){} // doesnt work!
	
	// var unsetted_var int
	// if(unsetted_var){} // doesnt work indepedent of the value! ONLY BOOLEAN CAN BE USED THIS WAY
	
	myStr := "some text"
	// if(myStr){} // doesnt work!
	// if(!myStr){} // doesnt work!
	if(myStr == ""){
		fmt.Println("myStr is a empty string")	
	}
	var pointer *int = &number1
	// if(pointer){} // doesnt work!
	// if(!pointer){} // doesnt work!
	if(pointer == nil){
		fmt.Println("pointer is nil")	
	}

	// other checking
	// number1 = 3
	// floatNumber := 3.0
	// if(number1 == floatNumber){} // WE CAN'T COMPARE INT WITH FLOAT
}
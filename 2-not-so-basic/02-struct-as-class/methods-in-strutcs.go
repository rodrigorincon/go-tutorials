package main

import(
	"fmt"
)

type Person struct{
	Name string
	Age int

	// presentation yourself with your dialect
	Presentation func(personName string)string // the attribute is a function, but we can't write the body inside struct. We can only define the header
											   // it's used for dependecy injection and add custom behavior
	//Presentation2 func(personName string)string{ fmt.Println("helo world") } // doesnt work!! 
}

// add a method in a struct
func (p Person) walk(meters int){
	fmt.Println(p.Name, "walk", meters, "meters")
}
func (p Person) birthdayByValue(){ // this method is useless. The struct is a copy, passed as copy by the function as the params
	p.Age += 1 					   // it's very good to avoid inappropriate methods change the values 
}
func (p *Person) birthdayByRef(){ // this method works because it receives the struct as reference, so we can change the attributes
	p.Age += 1
}
func (p Person) getAge()int{
	return p.Age
}

// REAL structure of a function
// func PARAM PASSED BEFORE THE NAME name(PARAM type)RETURN_TYPE {}
// the struct the "calls" the method is also a param

func main(){
	// we can't call these methods outside a struct
	// birthday() // doesnt work!
	// walk(3) // doesnt work!
	// age := getAge() // doesnt work!
	
	jhon := Person{
		Name: "Jhon", 
		Age: 10, 
		Presentation: func(otherPersonName string)string{
			return "Hello "+otherPersonName+". My name is "+"Jhon"+". Nice to meet you." // so formal
		},
	}
	mary := Person{
		Name: "Mary", 
		Age: 20, 
		Presentation: func(otherPersonName string)string{
			return "What's up man! I'm Mary." // not so formal
		},
	}
	fmt.Println(jhon.Name, jhon.Age)
	jhon.walk(4)
	fmt.Println("age before birthday", jhon.getAge())
	jhon.birthdayByValue()
	fmt.Println("age after birthday", jhon.getAge())
	jhon.birthdayByRef()
	fmt.Println("age after birthday", jhon.getAge())

	fmt.Println( jhon.Presentation("joseph") )
	fmt.Println( mary.Presentation("joseph") )
}
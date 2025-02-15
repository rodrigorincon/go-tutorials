package main

import "fmt"

type Person struct{
	name string
	age int32
}

func (p Person) talk(text string){ // add a method in the struct
	fmt.Println(p.name, "says: ", text) // we can access p attributes
}

func (p Person) cantUpdateAttrs(newValue string){ // we can read, but we can't write because p is a copy, not the real var
	p.name = newValue
}
func (p *Person) nowCanUpdateAttrs(newValue string){ // when we pass p as a pointer we can update it
	p.name = newValue
}

func (p Person) myMemoryAddress(){ // printing the memory location of the copy
	fmt.Printf("memory loc of the regular function: %p \n", &p)
}
func (p *Person) myMemoryAddressPointer(){ // printing the memory location of the original var
	fmt.Printf("memory loc of the pointered function: %p \n", p)
}

func main(){
	// when add functions in the struct scope it works like a class
	myVar := Person{"joao", 10}
	myVar.talk("hello, nice to meet you")

	fmt.Println("original name:", myVar.name)
	myVar.cantUpdateAttrs("adamastor")
	fmt.Println("after call function:", myVar.name)
	
	myVar.nowCanUpdateAttrs("adamastor")
	fmt.Println("after call function with pointer:", myVar.name)

	fmt.Println("-----------")

	fmt.Printf("my original memory loc: %p \n", &myVar)
	myVar.myMemoryAddress()
	myVar.myMemoryAddressPointer()
}
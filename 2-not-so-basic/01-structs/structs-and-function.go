package main

import "fmt"

type Person struct{
	name string
	age int32
}

func myFunc(p Person){ // copies the struct, changes here will not be affect other parts
	p.name = "func name"
}
func myFuncPointer(p *Person){
	p.name = "func name"
}

func main(){
	myVar := Person{"joao", 10}
	fmt.Println("before call function:", myVar.name) // joao
	myFunc(myVar)
	fmt.Println("after call function:", myVar.name) // joao
	myFuncPointer(&myVar)
	fmt.Println("after call function with pointer:", myVar.name) // func name
}
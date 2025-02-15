package main

import(
	"fmt"
)

type Person struct{
	Name string // is exported
	age int     // not exported
}
func (p Person) walk(){ // not exported
	fmt.Println(p.Name, "is walking")
}
func (p Person) Read(){ // is exported
	fmt.Println(p.Name, "is reading")
}

type persnoNotExported struct{
	Name string
	age int
}

const constNotExported = 4
const ConstExported = 5

var varNotExported int = 10
var VarExported int = 20

func funcNotExported(){}
func FuncExported(){}

func main(){
	// all vars, structs, objects and functions created are accessible inside your package!!!
	// so we can access all of them here

	// but only who starts with capitalized letter is exportable
	// we can only access capitalized vars, structs, objects and functions outside the package

	// in our package we can do:
	// package.ConstExported
	// package.VarExported
	// package.FuncExported()
	// p := package.Person{Name: ""}
	// p.Name
	// p.Read
	// OBS: WE CAN'T DO p.age! age IS A PRIVATE ATTRIBUTE
	// OBS: WE CAN'T DO p.walk! walk IS A PRIVATE METHOD
}
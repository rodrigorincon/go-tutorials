package main

import "fmt"

var globalVar = 10
var someVar = "global var"

// global var can be accessed in any place OF THIS PACKAGE. Global var can't be accessed outsde your package
func method(){
	fmt.Println("global var in the method:", globalVar)
}

func changeGlobalVar(){
	globalVar = 20
}

func main(){
	fmt.Println("global var in the main:", globalVar)
	method()

	// changing the globalVar
	changeGlobalVar()
	fmt.Println("global var in the main:", globalVar)
	method()

	// when we have a global and local vars with same name, the preference is the local
	someVar = "local var"
	fmt.Println(someVar)

	// ONLY USE GLOBAL VAR AS CONSTANT
}
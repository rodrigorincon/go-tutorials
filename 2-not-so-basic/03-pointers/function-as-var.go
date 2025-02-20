package main

import(
	"fmt"
	"reflect"
)

func myFunction(){
	fmt.Println("my function")
}

// function that returns a function
func returnFunc() func(param int)int{
	// the original function doesnt receive params and returns a function
	// the returned function receives a n int param and returns an int

	return func (value int)int{
		return value*2
	}
}

func receiveFuncAsParam(funcParam func(int)int){
	fmt.Println("Executing the param that I receive: ", funcParam(2))
}

func main(){
	// storing a function in a var
	varWithfunc := myFunction
	varWithfunc()
	fmt.Println("the type of this var is ", reflect.TypeOf(varWithfunc)) // type if func()
	fmt.Println("the var value is ", varWithfunc) // prints the function address

	// we can define a var as func type
	var funcVar func() = myFunction
	fmt.Println("the type of this var is ", reflect.TypeOf(funcVar)) // type if func()
	fmt.Println("the var value is ", funcVar)
	funcVar()

	// var can store a anonymous function
	anonymousFunc := func(param int)int {
		return param*2
	}
	fmt.Println("calling the anonymous func ", anonymousFunc(4))

	// a function can returns another function
	returnedVal := returnFunc()
	fmt.Println("calling the function returned ", returnedVal(3))

	// a function can receive another function as params
	receiveFuncAsParam( returnedVal )
}
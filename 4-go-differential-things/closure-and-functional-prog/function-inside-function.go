package main

import(
	"fmt"
)


func myFunction(){
	
	// func myFuncInside(){} // it doest work!!
	
	myFuncInside := func (p1 int)int{ // we can define a function inside a function as a var or running directly
		fmt.Println("I entered in myFunction->myFuncInside")
		return p1+1
	}
	fmt.Println( myFuncInside(1) )

	// running directly
	func (p1 int){ // we can define a function inside a function as a var or running directly
		fmt.Println("I entered in myFunction->anonymous function with param", p1)
	}(3)
}

func main(){
	myFunction()

	// closure is a anonymous function (internal) that access vars and inner functions of the function it's inside (external)
	counter := 1
	innerFunc := func(){
		fmt.Println("I'm inner funcion")
	}

	closureFun := func(param int) {
        fmt.Println("I'm in a closure. The external var is", counter)
        counter += param
        innerFunc()
    }
    // repair closure is identical to myFuncInside, excpet myFuncInside doesnt access external data. If it does myFuncInside will be a closure too

    fmt.Println("before call closure ", counter)
    closureFun(3)
    fmt.Println("after call closure ", counter)
}
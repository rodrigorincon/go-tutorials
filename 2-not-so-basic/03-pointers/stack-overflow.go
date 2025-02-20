package main

import(
	"fmt"
)

const SIZE_VAR int64 = 8

// it'll cause a stack overflow
func myFunction(value int64){
	value += 1
	fmt.Println("total space spent with vars: ", value*SIZE_VAR )
	myFunction(value)
}

func main(){
	var counter int64 = 2
	fmt.Println("total space spent with vars: ", counter*SIZE_VAR )
	myFunction(counter)
	// the stack memory limit is 1 billion bytes. It depends of the compiler and the OS
}

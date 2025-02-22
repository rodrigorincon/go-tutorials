package main

import( // import searches inside GOPATH directory, so our packages must be inside there
	"fmt"
	"my_first_package" // to work we need to copy the my_first_package folder into the GOPATH directory
)

func MyMethod(){
	fmt.Println("I'm a local function with same name as a function imported")
}

// just for show the created package, we keep the my_first_package directory here in the same directory, but this files are not used, only the copy in GOPATH
func main(){
	fmt.Println("I'm in the main method")
	my_first_package.MyMethod()

	MyMethod()

	fmt.Println("var imported: ", my_first_package.MyFirstVar)
}
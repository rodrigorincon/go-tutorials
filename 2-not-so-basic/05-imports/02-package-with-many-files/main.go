package main

import( // import searches inside GOPATH directory, so our packages must be inside there
	"fmt"
	"my_second_package" // to work we need to copy the my_second_package folder into the GOPATH directory
)

func MyMethod(){
	fmt.Println("I'm a local function with same name as a function imported")
}

// just for show the created package, we keep the my_second_package directory here in the same directory, but this files are not used, only the copy in GOPATH
func main(){
	fmt.Println("I'm in the main method")
	my_second_package.MyMethod()
	my_second_package.MyMethod2()

	MyMethod()

	fmt.Println("vars imported: ", my_second_package.MyFirstVar, "and", my_second_package.MySecondVar)

	// testing if a var created in file1 can be used in file2 in the same package
	my_second_package.UsingSameVar()
}
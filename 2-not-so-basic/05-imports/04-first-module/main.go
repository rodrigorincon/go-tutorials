package main

import (
	"fmt"
	"github.com/google/uuid" // to import libs from a git repo you need to create a mod
)

// to create a module and install the dependencies (external libraries) runs these 2 commands
// go mod init project_name // project_name is your project
// go get github.com/google/uuid // add the external lib in the mod file
func main(){
	myUuid := uuid.New() // the library name is the last folder path
	fmt.Println("my generated uuid:", myUuid)
}
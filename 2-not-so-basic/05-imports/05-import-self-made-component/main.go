package main

import (
	"fmt"
	"self_made_component/my_package" // NAME_USED_TO_NAME_MODULE_IN_GO.MOD/package
)

// to install a self-made component don't published in the internet (e.g. dividing your project in modules) you need to create a module with "go mod init"
// to import use MODULE_NAME/package_name where MODULE_NAM is the name used in go.mod and in "go mod init" command
// the packages needs to be in the same folder as the go.mod
func main(){
	my_package.MyMethod()
	fmt.Println("result returned by my component:", my_package.MyOtherMethod())
}
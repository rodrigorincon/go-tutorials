package my_first_package // the package must have the same name as the directory. It's required! 
// The compiler doesn't find the package to import if the package and the directory when it's there have different names

import(
	"fmt"
)

var MyFirstVar int = 5

func MyMethod(){
	fmt.Println("I'm in the first file method")
}

package my_second_package // we can use many files with same package, the import will wrap everything and will looks like all files were just one

import(
	"fmt"
)

var MyFirstVar int = 5

func MyMethod(){
	fmt.Println("I'm in the first file method")
}

// it doesnt work! We can't have 2 vars/struct/functions with same name in a package, even in different files
//func SameMethod(){ fmt.Println("This is the methdos replicated in most than 1 file. This is the version of first file") }

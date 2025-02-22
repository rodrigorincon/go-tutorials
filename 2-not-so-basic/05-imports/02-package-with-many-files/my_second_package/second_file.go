package my_second_package // we can use many files with same package, the import will wrap everything and will looks like all files were just one

import(
	"fmt"
)

var MySecondVar int = 10

func MyMethod2(){
	fmt.Println("I'm in the second file method")
}

// it doesnt work! We can't have 2 vars/struct/functions with same name in a package, even in different files
//func SameMethod(){ fmt.Println("This is the methdos replicated in most than 1 file. This is the version of second file") }

func UsingSameVar(){ // use a var created in the first_file, but both are in same package, I can access it here
	fmt.Println("I'm in the second file accessing a var created in the first file: MyFirstVar = ", MyFirstVar)
	fmt.Println("the same works for functions: calling MyMethod:")
	MyMethod()
}
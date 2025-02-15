package main

import "fmt"

// TYPE creates a new var type, in this case, a complex var with many attributes beause is a struct
type MyStruct struct{
	attr1, attr2 int
	attr3 string
}

func main(){
	var myVar MyStruct = MyStruct{1,2,"hello"} // set the vars in the order they're declared
	fmt.Println(myVar)
	fmt.Println("attr1:", myVar.attr1)
	fmt.Println("attr2:", myVar.attr2)
	fmt.Println("attr3:", myVar.attr3)

	myVar2 := MyStruct{1,2,"hello"}
	fmt.Println(myVar2)
	fmt.Println("attr1:", myVar2.attr1)
	fmt.Println("attr2:", myVar2.attr2)
	fmt.Println("attr3:", myVar2.attr3)

	// we can set the vars indicating specifically the attributes
	myVar3 := MyStruct{
		attr3: "hello",
		attr1: 1,
		attr2: 2, // we need to add the comma in the last attribute!!
	} // in this case, is not necessary follow the order in the struct declaration
	fmt.Println(myVar3)
	fmt.Println("attr1:", myVar3.attr1)
	fmt.Println("attr2:", myVar3.attr2)
	fmt.Println("attr3:", myVar3.attr3)
}
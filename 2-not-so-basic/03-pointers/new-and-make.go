package main

import(
	"fmt"
)

type Person struct{
	name string
	age int
	useGlasses bool
	mother *Person
}

// use new for structs and primitive types and make for slices, maps and channels
// new allocate memory and returns the pointer, make allocates and set the capacity and returns the object itself
func main(){
	// allocating memory with new
	var pointerInt *int = new(int)
	fmt.Println("address: ", pointerInt, "value: ", *pointerInt )
	var pointerStruct *Person = new(Person)
	fmt.Println("address:", &pointerStruct, "name:", pointerStruct.name, "age:", pointerStruct.age, "use useGlasses?", pointerStruct.useGlasses, "mother:", pointerStruct.mother )
	pointerSlice := new([]int) // creates a pointer for a slice with value = [] and length = 0
	fmt.Println("address: ", &pointerSlice, "value: ", *pointerSlice, "length:", len(*pointerSlice) ) 
	pointerSlice2 := new([2]int) // creates a pointer for a slice with value = [0,0] and length = 2
	fmt.Println("address: ", &pointerSlice2, "value: ", *pointerSlice2, "length:", len(*pointerSlice2) ) 

	// new allocates memory zeroed (int = 0, bool = false, pointer = nil and string = "")
	// receives the type of the var and returns a pointer for it
	// used for structs or primitive vars (int, bool, string...)
	fmt.Println("--------------")

	// allocating memory with make
	mySlice := make([]int, 2)
	fmt.Println("value: ", mySlice, "length:", len(mySlice), "first value:", mySlice[0] ) // prints [0 0] length: 2 first value: 0
	mySlice2 := make([]int, 2, 5)
	fmt.Println("value: ", mySlice2, "length:", len(mySlice2), "first value:", mySlice2[0] ) // prints [0 0] length: 2 first value: 0

	// make allocates memory and initialize the maximum capacity you need to allocate more
	// receives the type of the var, the length and the capacity and returns the object (not the pointer)
	// make doesnt work with pointers
	// make only works with slice, map, or channel 

	// mySlice := make([]int) // it doesnt work! make must receive 2 or 3 params
	// var myInt int = make(int, 2) // it doesnt work! MAKE ONLY WORKS WITH slice, map, or channel 
	// var myPerson Person = make(Person) // it doesnt work! MAKE ONLY WORKS WITH slice, map, or channel 

	// ABOUT MAKE ARGS ---------------------------
	// second arg of make is the official length of the object (size) and the third is the true size (capacity)
	// even you print it and showing just [0,0] and can't access mySlice2[2], the true size is 5. 
	// It's used to avoid new allocations when you do "append", what'll use the already allocated 3 remaining spaces
}

package main

import "fmt"

type Pessoa struct {
	name string
	age int8
}

func arrays(){
	// initializing and reading array
	var arr1 [2]string // define an array with 2 spaces
	arr2 := []int {1,2,3,4,5,6} // define an array and set it
	arr3 := [...]int {1,2,3} // using [...] MUST define the values in the same line and the compiler will define the size
	fmt.Println(arr1) // prints empty array
	fmt.Println( len(arr1) ) // return 2 because the array was initializes with 2 spaces 
	fmt.Println(arr3)

	fmt.Println(arr2)
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(len(arr2))

	// nil arrays
	var arr4 []uint8 = nil
	fmt.Println(arr4) // prints empty array
	fmt.Println( len(arr4) ) // prints 0

	// arr5 := []int // we need to initialize this way
	arr5 := []int(arr2) // we can initialize an array using this syntax by casting
	fmt.Println(arr2, arr5)

	// arr4 := [1,2,3,4,5] // arrays can't be initialized that way
	// var arr4 = [len(arr2)]int // array size must be constant in the compile moment, it can't be a var. For it, use malloc 

	// we can compare arrays
	arr6 := [2]int {1,2}
	arr7 := [2]int {1,2}
	fmt.Println( arr6 == arr7) // returns true
	arr7[1] = 3
	fmt.Println( arr6 == arr7) // returns false

	arr8 := [3]int {1,2,3}
	fmt.Println(arr8)
	// fmt.Println( arr6 == arr8) // compare arrays of size 2 and 3 breaks


	// comparing array of structs
	person1 := Pessoa{"ana", 1}
	var person2 Pessoa = Pessoa{"maria", 2}

	people1 := [2]Pessoa { person1, person2 }
	people2 := [2]Pessoa { Pessoa{name: "ana", age: 1}, Pessoa{name: "maria", age: 2} }
	fmt.Println( people1 == people2 ) // returns true

	people2[1].age = 3
	fmt.Println( people1 == people2 ) // returns false
}

func matrix(){
	var board [2][3]int
	for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = i + j
        }
    }
    fmt.Println(board)

    // board2 := []int {[1,2],[3,4]} // this syntax doesnt work!!
    // board2 := [][]int // this syntax doesnt work!!
}

func main(){
	fmt.Println("array")
	arrays()
	fmt.Println("matrix")
	matrix()
}
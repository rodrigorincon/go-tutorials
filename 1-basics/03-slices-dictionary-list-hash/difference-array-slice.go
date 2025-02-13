package main

import(
	"fmt"
	"reflect"
)

func main(){
	// difference between array and slice

	// array is compound var ([]T), FIXED size and all inner vars are of same type
	// examples:
	var arr1 [2]string
	arr2 := [...]int {1,2,3}
	arr3 := [6]int {1,2,3,4,5,6}
	fmt.Println( reflect.TypeOf(arr1).Kind() )
	fmt.Println( reflect.TypeOf(arr2).Kind() )
	fmt.Println( reflect.TypeOf(arr3).Kind() )

	// slice is compound var ([]T), VARIABLE size and all inner vars are of same type
	// examples:
	slice1 := []int {1,2,3,4,5,6}
	var slice2 []uint8
	slice3 := []int(slice1)
	slice4 := make([]int, 6)
	fmt.Println( reflect.TypeOf(slice1).Kind() )
	fmt.Println( reflect.TypeOf(slice2).Kind() )
	fmt.Println( reflect.TypeOf(slice3).Kind() )
	fmt.Println( reflect.TypeOf(slice4).Kind() )

	// we can't compare array with slice
	//ex: arr3 is [6]int and slice1 is []int even slice1 has 6 lengh

	// slices can be used in some methods that array can't
	slice5 := append(slice1, 7)
	fmt.Println(slice1, slice5)
	// arr2 = append(slice1, 7) // didnt work

	copiedSlice := make([]int, len(slice5))
    copy(copiedSlice, slice5)
    fmt.Println("copy:", copiedSlice)
    // var arr4 [6]int
    // copy(arr4, arr3) // didnt work
}
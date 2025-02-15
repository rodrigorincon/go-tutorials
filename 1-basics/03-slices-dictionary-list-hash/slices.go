package main

import(
	"fmt"
	"slices"
)

func minMax(){
	myList := []int{7, 9, 8, 1, 2, 3}

	max := slices.Max(myList)
	fmt.Println("Max in Slice =", max)

	min := slices.Min(myList)
	fmt.Println("Min in Slice =", min)
}

func binarySearch(){
	myList := []int{1, 2, 3, 4, 5, 6}
	index, found := slices.BinarySearch(myList, 2)
	fmt.Println("2 is in the array?", found, "in the index", index)
	index, found = slices.BinarySearch(myList, 41)
	fmt.Println("41 is in the array?", found, "in the index", index) // if not found show the index where it fits in the orderd array
}

func contains(){
	numbers := []int{0, 1, 2, 3}
	fmt.Println(slices.Contains(numbers, 2))
	fmt.Println(slices.Contains(numbers, 4))
}

func containsFunc(){
	numbers := []int{0, 2, -10, 4}
	hasNegative := slices.ContainsFunc(numbers, func(n int) bool {
		return n < 0
	})
	fmt.Println("Has a negative:", hasNegative)

	hasOdd := slices.ContainsFunc(numbers, func(n int) bool {
		return n%2 != 0
	})
	fmt.Println("Has an odd number:", hasOdd)
}

func find(){
	numbers := []int{0, 2, -10, 4, -4, 8}
	index:= slices.IndexFunc(numbers, func(n int)bool{
		return n < 0
	})
	fmt.Println("first negative numebr:", numbers[index])
}

func uniq(){
	seq := []int{0, 1, 1, 2, 3, 3, 5, 8}
	seq = slices.Compact(seq)
	fmt.Println("list with all repeated vars next", seq)

	seq = []int{0, 1, 2, 3, 5, 1, 8, 5}
	seq = slices.Compact(seq)
	fmt.Println("list but repeated vars are not next", seq)
}

func main(){
	fmt.Println("---- min and max - only in slices")
	minMax()
	fmt.Println("---- run a binary search in a sorted slice")
	binarySearch()
	fmt.Println("---- contains a value")
	contains()
	fmt.Println("---- check if exist some value that pass in the condition")
	containsFunc() // like "any?" in ruby
	fmt.Println("---- remove repeated objects - but only if they're next")
	uniq()
	fmt.Println("---- find the first var the pass in the criteria")
	find()
}
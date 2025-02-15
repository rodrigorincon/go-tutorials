package main

import "fmt"

func main(){
	// range array and slices
    arrList := []int{1,2,3,4,5}
    for _, value := range arrList {
        fmt.Println(value)
    }
    fmt.Println("---")

    // range string
    for _, letter := range "abcdef" {
        fmt.Println("letter", string(letter), "number ascii for the letter ", letter)
    }
    fmt.Println("---")

    // range string byte-to-byte
    const byteChar = "日本語"
    for i := 0; i < len(byteChar); i++ { // 3 bytes for each symbol
        fmt.Println(byteChar[i])
    }
    fmt.Println("---")

    // range map
    m := map[string]int{"one":   1,"two":   2,"three": 3,}
    for key, value := range m {
        fmt.Println(key, value)
    }
    fmt.Println("---")

    // RANGE CREATES A COPY OF THE OBJECTS, SO YOU CANT UPDATE THEM IN THE LOOP
    arrList = []int{1, 2, 3}
    for _, n := range arrList {
        n += 1 // change to [2,3,4]
    }
    fmt.Println(arrList) // prints [1,2,3]
    fmt.Println("---")

    val1 := 1
    val2 := 2
    val3 := 3
    pointerList := [3]*int{&val1,&val2,&val3}
    for _, pointer := range pointerList {
        *pointer += 1 // change to [2,3,4]
    }
    fmt.Println(*pointerList[0], *pointerList[1], *pointerList[2]) // prints [2,3,4]
    fmt.Println("---")

    // receiving only 1 var from range
    primes := []int{2, 3, 5, 7}
    for p := range primes { // if we receive only 1 var from range, it'll be the index (the first), so we'll lost the value
        fmt.Println(p) // print [0,1,2,3]
    }
}
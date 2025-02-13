package main

import (
    "fmt"
)

func main() {
    // var number8 uint8 = nil // We can't initialize a var as null
    // notTypedVar := nil //  We can't initialize a var as null

    var numberArray []uint8 = nil // we can initialize an array as null
    fmt.Println( numberArray, len(numberArray) ) // [], 0

    fmt.Println(numberArray == nil) // true
    //fmt.Println(numberArray == []) // we can't compare null with something
    
    var typedVar uint32 = 60
    notTypedVar := 70
    // typedVar = nil // we can't change a already existed var to nil
    // notTypedVar = nil // we can't change a already existed var to nil

    fmt.Println(typedVar)
    fmt.Println(notTypedVar)

    // null in pointer
    var someInt int = 5
    var pointer *int
    fmt.Println(pointer) // initialized as null
    pointer = &someInt
    fmt.Println(pointer) // some address
    pointer = nil
    fmt.Println(pointer) // null
}
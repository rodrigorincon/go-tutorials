package main

import (
    "fmt"
    "encoding/binary"
)

func main() {
    var number8 uint8 = 65
    numberRune := rune(number8) // casting to rune (int32)
    numberStr := string(number8) // casting to string
    
    fmt.Println(number8)
    fmt.Println( numberStr ) // read as ascii
    fmt.Println( numberRune ) // read as ascii
    // we can't compare different types (numberRune == number8 don't compile) 
    
    // casting int32 to uint8[]
    var number32 uint32 = 65
    buffer := []uint8 {0,0,0,0}
    binary.BigEndian.PutUint32(buffer, number32)
    fmt.Println(buffer)
    number32 = 31_415_926
    binary.BigEndian.PutUint32(buffer, number32)
    fmt.Println(buffer)

    // calculating the parts of array to check if iit's the same of original
    var sum_values uint32 = uint32(buffer[0])*256*256*256 + uint32(buffer[1])*256*256 + uint32(buffer[2])*256 + uint32(buffer[3])
    fmt.Println(sum_values == number32)

    // casting string to rune[]
    newByteArray := []rune("hello world")
    fmt.Println(newByteArray)
}
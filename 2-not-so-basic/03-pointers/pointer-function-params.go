package main

import (
    "fmt"
)

func main() {
    var age int = 10
    fmt.Println("address in main func:", &age, "value:", age)
    age = multiply(age, 3)
    fmt.Println("address in main func:", &age, "value:", age)

    referenceMultiply(&age, 3)
    fmt.Println("address in main func:", &age, "value:", age) // age is changed in the method and it's reflected in main, becuase it's the same memory space
}

func multiply(value int, multiplier int) int {
    fmt.Println("address in another func:", &value, "value:", value)
    return value * multiplier
}

func referenceMultiply(intPointer *int, multiplier int) {
    result := (*intPointer) * multiplier // (*intPointer) converts to pointer to int, or better, gets the pointer intPointer
    *intPointer = result // *intPointer informs that we're reading/writing in the memory destination, so we're writing the result in memory pointed by intPointer, 
                         //not writing in the var intPointer
                         // BUT if the pointer type is a struct I dont need to use the *, I can change the value as the pointer is the true object
    fmt.Println("address in reference func:", intPointer, "value:", *intPointer)
}
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var my_float float32 = 1.2345
    fmt.Println("my_float", my_float)
    my_float = -10.54321
    fmt.Println("my_float", my_float)

    /* float types
    float32 from 1,5 x 10−45 to ±3,4 x 1038 (32 bits) - 6 to 9 precision digits
    float64 from ±5.0 × 10−324 para ±1.7 × 10308 (64 bits) - 15 to 17 precision digits

    DOESNT EXIST UNSIGNED FLOAT
    */

    // IMPORTANT: doesnt exist float alias
    // var just_float float = 5.1 // THIS LINE WILL BREAK

    ////// explicity and implicity declaration
    var explicity_float float32 = 5
    fmt.Println("explicity_float", explicity_float)

    implicity_folat := 3.14 // the size will be float64 because my machine is 64
    fmt.Println("implicity_folat", implicity_folat)
    fmt.Println("the number of bytes is", unsafe.Sizeof(implicity_folat) )


    /////// FLOATS AND INTEGERS
    var integerNum int32 = 3
    var floatNumber float32 = 3.0
    // floatNumber = integerNum // WE CAN'T PASS INT VALUE TO FLOAT
    floatNumber = float32(integerNum) // we need to convert to work
    fmt.Println("converted value to floar", floatNumber)
    // if(number1 == floatNumber){} // WE CAN'T COMPARE INT WITH FLOAT

    /////// COMPLEX NUMBERS
    var complex1 complex64 = 10 + 23i // imaginary float32
    var complex2 complex128 = 12i // imaginary float64
    complex3 := complex(10, 3) // will be 32 or 64 depending the architecure. The first value is real and second the imaginary
    complex4 := 4i // // will be 32 or 64 depending the architecure
    fmt.Println("complex1", complex1)
    fmt.Println("complex2", complex2)
    fmt.Println("complex3", complex3)
    fmt.Println("complex4", complex4)

    /////// boolean
    bool_var := true
    bool_var2 := false
    fmt.Println("bool_var", bool_var)
    fmt.Println("bool_var2", bool_var2)

    var bool_var3 bool = true
    fmt.Println("bool_var3", bool_var3)
    var bool_var4 bool // starts with FALSE
    fmt.Println("bool_var4", bool_var4)
    // var bool_var5 boolean = true // IMPORTANT: BOOLEAN TYPE DOESNT EXIST, JUST BOOL
}
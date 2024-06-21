package main

import (
    "fmt"
    "unsafe"
)

func main() {
    //// positive and negative values
    var my_int int
    my_int = 5
    fmt.Println(my_int)
    my_int = -10
    fmt.Println(my_int)
    int_text := 1234
    fmt.Println( int_text )

    ///// more readable way to write a number
    int_text2 := 1_000_000
    fmt.Println( int_text2 )

    //// operations
    my_number := 10
    my_number += 5
    fmt.Println(my_number) // 15
    my_number -= 7
    fmt.Println(my_number) // 8
    my_number *= 2
    fmt.Println(my_number) // 16
    my_number /= 4
    fmt.Println(my_number) // 4
    my_number %= 3
    fmt.Println(my_number) // 1
    fmt.Println(9%6) // 3
    other_number := 6
    my_number += other_number
    fmt.Println(my_number) // 7

    /* int types
    int8    from -128 to 127 (8 bits)
    uint8   from 0 to 255    (8 bits)
    int16   from -32.768 to 32.767 (16 bits)
    uint16  from 0 to 65.535       (16 bits)
    int32   from -2.147.483.648 to 2.147.483.647 (32 bits)
    uint32  from 0 to 4.294.967.295              (32 bits)
    int64   from -9.223.372.036.854.775.808 to 9.223.372.036.854.775.807 (64 bits)
    uint64  from 0 to 18.446.744.073.709.551.615                         (64 bits)
    uintptr equals to uint (32 or 64 depending your architecture) used to do pointers
    */

    ///// size and types

    // IMPORTANT: int is an alias for or int32 or int64, depending the acrhitecture of the machine when it was compiled
    var myint32 int32 = 5
    var myint64 int64 = 5
    fmt.Println( unsafe.Sizeof(my_int) ) // as my machine is 64bits, int is int64
    fmt.Println( unsafe.Sizeof(myint32) )
    fmt.Println( unsafe.Sizeof(myint64) )

    ///// overflowing the limits

    var positive_number uint8 = 5
    fmt.Println(positive_number)
    positive_number = positive_number - 10
    fmt.Println(positive_number) // returns 251
    // instead raise an error, it goes to greater value and continues to subtract (subtract 5 until goes to 0, subtract more 1 and instead goes to -1, goes to greater value - 255 - en subtract more 4)

    var value64 int64 = 9223372036854775806
    value64 = value64 * 100
    fmt.Println(value64) // doesnt raise excpetion, just return a non sense number

    int64value := 4294967295
    fmt.Println(int8(int64value) )

    ///// alias
    var byte_number byte = 45 // byte is an alias to uint8 (just positives)
    fmt.Println(byte_number)

    var rune_number rune = 100 // rune is an alias to int32, usually is used to store numeracly Unicode characters
    fmt.Println(rune_number)
    fmt.Println("the unicode char of this rune is", string(rune_number))
    var regular_int int = 99
    fmt.Println("I can print the ascii value for each number, not just rune:", regular_int, string(regular_int))


    ///// converting
    fmt.Println( unsafe.Sizeof(123) ) // 8 bytes because my machine is 64bits
    fmt.Println( unsafe.Sizeof( int8(123) ) ) // convert number to int8, so the answer is 1 byte
    new_int := 123
    fmt.Println( unsafe.Sizeof(new_int) ) // 8 bytes because my machine is 64bits
    fmt.Println( unsafe.Sizeof( int8(new_int) ) ) // convert number to int8, so the answer is 1 byte
}
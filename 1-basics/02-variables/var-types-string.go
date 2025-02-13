package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    //// defining
    fmt.Println(`///// DEFINING`)
    my_char := "a"
    // var my_char char = "a" // IMPORTANT: "CHAR" TYPE DOESNT EXIST
    fmt.Println("my_char: ", my_char)
    fmt.Println("size of the char: ", len(my_char))
    
    var my_string string = "a"
    fmt.Println("my_string: ", my_string)
    my_string = "abcde"
    fmt.Println(my_string)
    fmt.Println("size of the string: ", len(my_string))
    
    //// type size
    fmt.Println(`///// TYPE SIZE`)
    my_char2 := "a"
    fmt.Println("number of bytes of a char: ", unsafe.Sizeof(my_char2) ) // 16 bytes even been 1 character
    my_string2 := "abcde"
    fmt.Println("number of bytes of a string: ", unsafe.Sizeof(my_string2) ) // also have 16 bytes even been much more characters

    /* The reason for string always have 16 bytes is String is a struct
    type StringHeader struct {
        Data uintptr -> 8 bytes for the pointer
        Len  int     -> 8 bytes for the attribute thar stores the size of the string (until where read with the pointer)
    }
    */

    ////// accessing specific positions
    fmt.Println(`///// ACCESSING SPECIFIC POSITIONS`)
    my_string2 = "abcde"
    fmt.Println("my_string2[1]: ", my_string2[1]) // print the uft8 number (explaining in the end of this section)
    fmt.Println("string(my_string2[1]", string(my_string2[1]) ) // print the string
    char_read := my_string2[2]
    fmt.Println("char_read: ", char_read) // print the uft8 number
    fmt.Println("string(char_read): ", string(char_read) ) // print the string
    // why print a single char shows its number?
    // because a string is an ARRAY OF BYTES (byte = uint8), so a string is an array of ints that the compiler knows that have tread as a text
    // when you separate a slice of it and process separatelly it thinks it's just a uint8
    // example:
    my_string2 = "abcde"
    char_read = my_string2[1] // b or 98
    some_number := char_read + 12 // 110
    fmt.Println("some_number: ", some_number) // 110

    ////// strings are "immutable"
    fmt.Println(`///// IMMUTABLE`)
    my_string3 := "abcde"
    fmt.Println("my_string3: ", my_string3, "in position", &my_string3)
    my_string3 = "12345" // by address we see it's not the same object. The memory address changes, a new string was created and the last is lost
    fmt.Println("my_string3: ", my_string3, "in position", &my_string3)
    //my_string3[1] = "a" // but I can't change just a piece

    //// iterating
    fmt.Println(`///// ITERATION`)
    for index, for_char := range "Hello" {
        fmt.Println("position ", index, " is ", string(for_char))
    }

    ////// single quotes (ALWAYS USE DOUBLE QUOTES FOR STRING OR JUST 1 CHAR, NEVER SINGLE QUOTES!!!)
    fmt.Println(`///// SINGLE QUOTES`)
    single_quote := 'a' // IMPORTANT: this will not stores the char, but the utf8 number
    // single quote is not string, is a utf8 format
    fmt.Println("single_quote: ", single_quote) // print int32
    fmt.Println("reflect.TypeOf(single_quote): ", reflect.TypeOf(single_quote) )

    // single_quote2 := 'abcde' // IMPORTANT: IT WILL BREAK
    // double quotes work ok
    double_quote := "a"
    double_quote2 := "abcde"
    fmt.Println("double_quote: ", double_quote)
    fmt.Println("double_quote2: ", double_quote2)

    ///// crasis
    crasis_string := `this is my string: "hello world"\nThis is a new line`
    quote_string := "this is my string: \"hello world\"\nThis is a new line"
    fmt.Println("crasis_string: ", crasis_string)
    fmt.Println("quote_string: ", quote_string)

    ///// UTF8
    utf8_string := "Hello, 世界" // GO uses utf8 encoding by default
    fmt.Println("utf8_string: ", utf8_string)
    fmt.Println("len(utf8_string): ", len(utf8_string) ) // each special character occupies 3 characters
}
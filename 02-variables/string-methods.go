package main

import (
    "fmt"
    "strings"
)

func reverse(original_string string) string { 
    final_str := []rune(original_string) // convert to an array of rune 
    for i, j := 0, len(final_str)-1; i < j; i, j = i+1, j-1 { // just iterate until the middle of the string
        // swap the letters of the string, 
        // like first with last and so on. 
        final_str[i], final_str[j] = final_str[j], final_str[i] 
    } 
    // return the reversed string. 
    return string(final_str) 
} 

func main() {
    // copy
    fmt.Println("////// COPY")
    my_str := "abcdef"
    fmt.Println("First string:", my_str, "In the position ", &my_str)
    copied := my_str
    fmt.Println("Copied string:", copied, "In the position ", &copied) // a new var in a new address
    copied += "ghij"
    fmt.Println("Copied changed string:", copied, "In the position ", &copied) // the same address
    fmt.Println("Unchanged Original string:", my_str)

    // rewriting the same value in the string
    my_str = "abcdef"
    fmt.Println("address ", &my_str) // address doesnt changes

    // get a substring 
    my_str = "abcdef"
    substr := my_str[1:3] // bc
    fmt.Println("substr ", substr, "address: ", &substr) 
    substr = "aaaa"
    fmt.Println("substr ", substr, "address: ", &substr) 
    fmt.Println("original is not changed ", my_str)

    // check if has a substring 
    my_str = "abcdef"
    contains := strings.Contains(my_str,"cde")
    fmt.Println("String contains the substring 'cde'? ", contains)
    contains2 := strings.Contains(my_str,"fgh")
    fmt.Println("String contains the substring 'fgh'? ", contains2)

    // get a substring index
    index := strings.Index(my_str, "cde")
    fmt.Println("Index position of the substring 'cde' ", index)
    index2 := strings.Index(my_str, "fgh")
    fmt.Println("Index position of the substring 'cde' ", index2) // when doesn't has the substring returns -1

    // invert
    my_str = "abcdef"
    reversed := reverse(my_str) // doesnt have a method to do it, so I need to create myself a method to do it
    fmt.Println("original: ", my_str)
    fmt.Println("reversed: ", reversed)

    // regex

    // replace substring

    // lower/uppercase

    // split

    // join arrays in a string

    // concat

    // Capitalize all first letter of a setence
}
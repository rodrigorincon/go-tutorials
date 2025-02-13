package main

import (
    "fmt"
    "strings"
)

func reverse(original_string string) string { 
    final_str := []rune(original_string) // convert to an array of rune (int32)
    for i, j := 0, len(final_str)-1; i < j; i, j = i+1, j-1 { // just iterate until the middle of the string
        // swap the letters of the string, 
        // like first with last and so on. 
        final_str[i], final_str[j] = final_str[j], final_str[i] 
    } 
    // return the reversed array and casting to string. 
    return string(final_str) 
} 

func main() {
    // copy
    fmt.Println("////// COPY")
    my_str := "abcdef"
    fmt.Println("First string:", my_str)
    copied := my_str
    fmt.Println("Copied string:", copied)
    copied += "ghij"
    fmt.Println("Copied changed string:", copied)
    fmt.Println("Unchanged Original string:", my_str)

    // get a substring 
    my_str = "abcdef"
    substr := my_str[1:3] // bc
    fmt.Println("substr ", substr) 

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

    // replace substring
    stringToChange := "your password is abc123 and your ecurity phrase is abc123"
    replacedString := strings.ReplaceAll(stringToChange, "abc123", "***")
    fmt.Println("replaced string: ", replacedString)

    stringToChange2 := "your password is abc123 and your ecurity phrase is hello"
    replacer := strings.NewReplacer("abc123", "MY-PASSWORD", "hello", "MY-PASSPHRASE")
    replacedString2 := replacer.Replace(stringToChange2)
    fmt.Println("replaced string 2: ", replacedString2)


    // lower/uppercase
    fmt.Println( strings.ToUpper("hello world") )
    fmt.Println( strings.ToLower("HELLO WORLD") )
    fmt.Println( strings.Title("hello world") )

    // split
    arrayOfStrings := strings.Split("hello-world","-")
    fmt.Println(arrayOfStrings[0])
    fmt.Println(arrayOfStrings[1])

    // join arrays in a string
    listOfWords := []string {"this", "is", "my", "test"}
    fmt.Println(strings.Join(listOfWords, "-"))

    // concat
    text1 := "first"
    text2 := "text"
    concatString1 := text1 + " " + text2
    fmt.Println(concatString1)

    text1 += text2
    fmt.Println( text1 )

    // .push method doesnt exist!!!
    // << is like C, move the bytes, doesnt concat
}
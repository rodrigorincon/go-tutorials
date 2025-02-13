package main

import (
    "fmt"
    "regexp"
)

func main() {
    // regex pattern to validate CPF
    cpfPattern, _ := regexp.Compile("([0-9]{2}[\\.]?[0-9]{3}[\\.]?[0-9]{3}[\\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\\.]?[0-9]{3}[\\.]?[0-9]{3}[-]?[0-9]{2})")

    // check if a string matches in a regex
    fmt.Println("MatchString tests")
    fmt.Println( cpfPattern.MatchString("hello world") ) // false
    fmt.Println( cpfPattern.MatchString("482.129.432-76") ) // true
    fmt.Println( cpfPattern.MatchString("482.129.432-76 hellow world") ) // true
    fmt.Println( cpfPattern.MatchString("482.129.helloworld432-76") ) // false

    // it worked with bytes too
    fmt.Println("Byte tests")
    byteArray := []byte("482.129.432-76")
    fmt.Println( cpfPattern.Match(byteArray) ) // true
    newByteArray := []byte("hello world")
    fmt.Println( cpfPattern.Match(newByteArray) ) // false

    // returns the part that matches for the regexp.
    fmt.Println("FindString tests")
    fmt.Println( cpfPattern.FindString("hello world") ) // returns empty string
    fmt.Println( cpfPattern.FindString("482.129.432-76") )
    fmt.Println( cpfPattern.FindString("482.129.432-76 hellow world") )
    fmt.Println( cpfPattern.FindString("482.129.helloworld432-76") ) // returns empty string

    // get the index when start and finish the match
    indexes := cpfPattern.FindStringIndex("482.129.432-76")
    fmt.Println("Start index", indexes[0], "End index", indexes[1])

    // return all parts that matches for the regex. The second param is the number of responses (a negative number returns all)
    fmt.Println("FindAllString tests")
    fmt.Println( cpfPattern.FindAllString("482.129.432-76 hello world 123.123.123-12", -1) )
    fmt.Println( cpfPattern.FindAllString("482.129.432-76 hello world 123.123.123-12 third option 321.654.987-54", 1) )

    // get the index when start and finish all the matches
    allIndexes := cpfPattern.FindAllStringSubmatchIndex("482.129.432-76 hello world 123.123.123-12 third option 321.654.987-54", -1)
    fmt.Println("Start index 1:", allIndexes[0][0], "End index 1:", allIndexes[0][1], 
                "Start index 2:", allIndexes[1][0], "End index 2:", allIndexes[1][1], 
                "Start index 3:", allIndexes[2][0], "End index 3:", allIndexes[2][1])

    // replace all matches by a new string
    fmt.Println("ReplaceAllString tests")
    fmt.Println( cpfPattern.ReplaceAllString("482.129.432-76 hello world 123.123.123-12 third option 321.654.987-54", "***") )
}
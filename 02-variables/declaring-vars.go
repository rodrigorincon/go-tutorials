package main

import "fmt"

func main() {
    ////// explicitly defining the type

    var my_var int // defining explicitly the type
    my_var = 5
    fmt.Println(my_var)

    var other_var int = 8 // defining explicitly the type and setting a value
    fmt.Println(other_var)

    // third_var uint = 5 // THIS DECLARATION IS WRONG! IF YOU DEFINE EXPLICITLY THE TYPE, YOU NEED TO USE "VAR"

    ////// implicitly defining the type

    third_var := 10 // defining implicitly the type
    fmt.Println(third_var)

    fourth_var := third_var // defining implicitly the type (will be the same type of the last one)
    fmt.Println(fourth_var)

    // var fifth_var := 15 // THIS DECLARATION IS WRONG! YOU CAN'T USE "VAR" WITH ":"
    var my_first_var = 123
    fmt.Println(my_first_var)
    
    // my_second_var = 1 // IMPORTANT! YOU CAN'T CREATE A VAR THAT WAY, ON IMPLICIT CREATION YOU NEED TO ADD ":", JUST AFTER YOU CAN CHANGE THE VALUE THIS WAY
    // my_first_var := 321 // IMPORTANT! WE CAN'T USE ":" TO UPDATE THE VALUE OF AN ALREADY EXISTED VAR
    
    ////// VARs can change the value

    my_var = 1
    other_var = 2
    third_var = 3
    fmt.Println(my_var)
    fmt.Println(other_var)
    fmt.Println(third_var)
    fmt.Println(fourth_var) // changing the "third_var" doesn't change the "fourth_var"

    // my_var := 3 // THIS DECLARATION IS WRONG! I CAN'T CHANGE A VALUE USING ":"
    // USE ":" ONLY WHEN CREATES A NEW VAR

    var explicity_var int
    explicity_var = 1
    //explicity_var = "lalala" // THIS DECLARATION IS WRONG! explicity_var IS INT, I CANT PASS A STRING FOR IT (STRONGLY TYPED)
    fmt.Println(explicity_var)
    implicity_var := 2
    //implicity_var = "lalala" // THIS DECLARATION IS WRONG! implicity_var IS INT, I CANT PASS A STRING FOR IT (STRONGLY TYPED)
    fmt.Println(implicity_var)

    ////// defining multiple vars at once

    var a, b int = 1, 2
    fmt.Println(a)
    fmt.Println(b)
    // var c, d := 3, 4 // THIS DECLARATION IS WRONG! I NEED TO DEFINE THE TYPE WHEN CREATES MANY VARs IN A SINGLE LINE
    var c, d int
    fmt.Println(c) // print 0
    fmt.Println(d) // print 0

    
    ////// defining booleans
    var bool = true
    fmt.Println(bool)

    ////// defininga var but printing without setting

    var new_var int
    fmt.Println(new_var) // print 0
    var new_var2 string
    fmt.Println(new_var2) // print ""
    fmt.Println(new_var2 == "") // confirming the var is an empty string and not null
    var new_var3 float64
    fmt.Println(new_var3) // print 0
}
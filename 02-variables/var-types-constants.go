package main

import ("fmt")


func main() {
    const my_const int = 30
    fmt.Println("explicit const: ", my_const)
    // my_const = 10 // IF I TRY TO CHANGE A CONST THE COMPILATION BREAKS

    // const my_const2 := 20 // CONST DOESNT HAVE IMPLICIT DECLARATION
    
    // WE CAN'T DEFINE A PARAMETER IN FUNCTION AS CONST
    // func my_function(const param1 int) // IT DOESN'T WORK
    // if you pass int or string, it's created a copy in the function, so don't worry
    // BUUUT, if you're passing pointers, structs, slices, channels or maps (complex data types), it can be changed and you can't to anything to protect the var
    // you could create a copy inside the method and work just with it, but it costs a lot of time and memory. Depdening the function and dta size, that's a terrible idea
    // SO THE ONLY OPTION IS BE ALERT!!!
}
package main

import "fmt"

func main(){
	// for doesnt need parentesis
	i := 0
	for i = 0; i < 5; i++ {
        fmt.Println(i)
    }
    fmt.Println("---")

    // for DOESNT WORK with parentesis
	// for(i = 0; i < 5; i++) {}

    // we can create the index var inside the for statement. In this case, the var doesnt exist inside the for block
	for index := 0; index < 5; index++ {
        fmt.Println(index)
    }
    fmt.Println("---")

	// infinite loop
	for { // if you dont pass anything in for statement, you'll create a infinite loop
        fmt.Println("Infinite Loop!")
        break
    }
	for(true) {
        fmt.Println("Infinite Loop 2!")
        break
    }
    fmt.Println("---")

    // range
    arrList := []int{1,2,3,4,5}
    for index, value := range arrList {
        fmt.Println("Index", index, "value", value)
    }
    fmt.Println("---")

    // break and continue
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        if i%2 == 0 {
            continue
        }
        fmt.Println("index: ", i) // must print 1 and 3
    }

    // GO DOESNT HAVE WHILE, but you can have the same behavior using for
    counter := 0
    for counter < 5 {  // create a for that behaves like a while
        fmt.Println(counter)
        counter++
    }


    /// SUMARIZING
    // for(initial point ; finish statement; increment) for with 3 parts (traditional)
    // for(finish statement) for with 1 part (inifite loop and while)
    // for(range) for with 1 part (with range)
}
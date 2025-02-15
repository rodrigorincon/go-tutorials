package main

import (
    "fmt"
    "unsafe"
)

type Pessoa struct {
    name string
    age int8
}

func main() {
    // printing the memory address
    var myInt int32 = 512
    fmt.Println("value:", myInt, "address:", &myInt) 

    var pessoa Pessoa = Pessoa{"joao da silva", 12}
    fmt.Println("name:", pessoa.name, "address:", &pessoa.name)
    fmt.Println("age:", pessoa.age, "address:", &pessoa.age)

    var pointer *string
    var otherVar int8 = 5
    fmt.Println("pointer address: ", &pointer)
    fmt.Println("next used address: ", &otherVar)
    fmt.Println("pointer value", pointer) // returns nil because the pointer is not pointed to any pace
    fmt.Println("pointer is null?", pointer == nil)
    
    // pointer size
    var pointer2 *int8
    fmt.Println("string pointer size:", unsafe.Sizeof(pointer) ) // pointer always have 8 bytes (or 4 if your computer is 32bits)
    fmt.Println("int8 pointer size:", unsafe.Sizeof(pointer2) )

    // poiting for some address
    pointer = &pessoa.name
    fmt.Println("pointer value", pointer)
    fmt.Println("value of pointer memory:", *pointer)

    // struct pointer
    var pointer3 *Pessoa
    pointer3 = &pessoa
    fmt.Println("value of struct pointer:", pointer3.name, pointer3.age)
    
    pointer3.name = "maria" // to update a struct attribute I dont need to use the *
    fmt.Println("new person name:", pessoa.name)

    // pointing to nil
    pointer3 = nil
    fmt.Println("pointer3 value:", pointer3) // if I do *pointer will throw an error
    if pointer3 == nil { // !pointer3 DOESNT WORK!!!!
        fmt.Println(" *pointer3 breaks because pointer is null and I cant read the value of none ")
    }
}
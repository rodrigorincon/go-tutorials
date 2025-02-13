package main

// THIS NEW FEATURE ONLY WORKS SINCE 1.24
import (
    "fmt"
    "runtime"
    "time"
    "weak"
)

type Pessoa struct {
    name string
}

func main() {
    var pessoa *Pessoa = &Pessoa{name: "Joao"} // pessoa is a strong pointer (points directly to memory)
    var strongPointer *Pessoa = pessoa // strongPointer is a strong pointer (points directly to memory)

    pessoa = nil

    runtime.GC()
    time.Sleep(1 * time.Second)

    fmt.Println(strongPointer == nil) // false
    fmt.Println(strongPointer.name)

    //////////////////////////////////////////

    var pessoa2 *Pessoa = &Pessoa{name: "Joao"} // pessoa is a strong pointer (points directly to memory)
    var weakPointer *Pessoa = weak.Make(pessoa2) // don't point directly, so if the original pointer is gone all others is gone too and the memory is clear
    pessoa2 = nil

    runtime.GC()
    time.Sleep(1 * time.Second)

    if weakPointer.Value() == nil {
        fmt.Println("Reference is gone", weakPointer == nil)
    }else{
        fmt.Println("Value is still there", weakPointer.name)
    }
}
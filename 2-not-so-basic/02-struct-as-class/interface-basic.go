package main

import (
    "fmt"
)

// interface with functions to be implemented by children
type geometry interface {
    area() float64
    perim() float64
}


type rect struct {
    width, height float64
}
// implementation of the interface methods
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
func (r rect) notRelatedMethod() {
    fmt.Println("I entered in notRelatedMethod")
}

func main() {
    // we can create a var with type interface and receive a concrete implementation of it 
    var inter geometry
    fmt.Println("my intarface var: ", inter)
    // inter.area() // it breaks because var is nil until we declare a concrete implementation

    inter = rect{3,4}
    fmt.Println("my intarface var: ", inter)
    // inter.notRelatedMethod() // it doesnt work! we need to casting first
    inter.(rect).notRelatedMethod()
    
    var inter2 geometry = rect{3,4}
    fmt.Println("my intarface2 var: ", inter2)
    // inter2.notRelatedMethod() // it doesnt work! we need to casting first
    inter2.(rect).notRelatedMethod()
}
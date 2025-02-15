package main

import (
    "fmt"
    "math"
)

// interface with functions to be implemented by children
type geometry interface {
    area() float64
    perim() float64
}

// structs
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}
// struct that DOESNT implement the interface
type line struct{
	length float64
}

// implementation of the interface methods
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius,2)
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// line DOESNT implement the interface methods because line is a pointer! To implements needs to remove the *
func (l *line) area() float64 {
    return 0
}
func (l *line) perim() float64 {
    return l.length
}

// function that works for any implementation of interface
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func detectGeometry(g geometry) {
	switch g.(type) {
	case rect:
		fmt.Println("I'm a rect")
	case circle:
		fmt.Println("I'm a circle")
	case geometry:
		fmt.Println("I'm a undefined geometry")
	default:
		fmt.Println("it's not a geometry")
	}
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    // l := line{length: 6}

	fmt.Println("---- measurements of a rectangle")
    measure(r)
    fmt.Println("---- measurements of a circle")
    measure(c)
    // fmt.Println("---- measurements of a line")
    // measure(l) // doesnt work because line doesnt implements correctly the functions

    detectGeometry(r)
    detectGeometry(c)

	var g geometry
    detectGeometry(g)
}
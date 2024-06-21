package main // ALL go files must start defining which package it belongs. In this case, this file belongs to package "main"

import "fmt" // importing fmt library to print in the screen

// the code always start by main function
func main() {
    myText := "hello world"
    fmt.Println(myText)
}

/* important notes:
Go is strongly typed
Compiled
High level language
Multi paradigm (imperative, functional, concurrent, OO (not so much) )
has garbage collector
Foccused in concurrency
Doesn't have exceptions, inheritance and overload (sobrecarga - métodos com mesmo nome mas params diferentes)
Memory safety (protect against programmer's errors, avoiding memory issues about memory)

to compile a code, run "go build FILE.GO" and after run the binary "./FILE"
Go doesn't have a built in interactive terminal command to run for fast tests (like node or irb), you need to download and install separately
to run tests, run "go test FILE_TEST.GO"
variables, constants and functions are camelCase. Public data have the first letter Capitalized and privated first letter lowerCased
*/
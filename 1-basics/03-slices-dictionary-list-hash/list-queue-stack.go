package main

import (
    "fmt"
)

func listArr(){
    fmt.Println("-------------- LIST")
    // create a slice (array) that behaves as a list. 
    // But every time we append or remove we duplicate the slice, so it's not optimized
    list := make([]int, 0) // initialize an empty list
    list = append(list, 0) // add value
    list = append(list, 1)
    list = append(list, 2)
    list = append(list, 3)
    list = append(list, 4)
    fmt.Println(list)
    // read a specific value by index
    fmt.Println( list[2] )
    // remove value by index
    index := 3
    list = append(list[:index], list[index+1:]...)
    fmt.Println(list)
}

func queueArr(){
    fmt.Println("-------------- QUEUE")
    // create a slice (array) that behaves as a queue. 
    // But every time we append or remove we duplicate the slice, so it's not optimized
    queue := make([]int, 0) // initialize an empty queue
    queue = append(queue, 0) // add value inserting in the end
    queue = append(queue, 1)
    queue = append(queue, 2)
    queue = append(queue, 3)
    queue = append(queue, 4)
    fmt.Println(queue)
    // read the first value of the queue
    value := queue[0]
    fmt.Println(value)
    // delete the first
    queue = queue[1:]
    fmt.Println(queue)
}

func stackArr(){
    fmt.Println("-------------- STACK")
    // create a slice (array) that behaves as a stack. 
    // But every time we append or remove we duplicate the slice, so it's not optimized
    stack := make([]int, 0) // initialize an empty stack
    stack = append(stack, 0) // add value inserting in the end
    stack = append(stack, 1)
    stack = append(stack, 2)
    stack = append(stack, 3)
    stack = append(stack, 4)
    fmt.Println(stack)
    // read the last value of the stack
    value := stack[ len(stack)-1 ]
    fmt.Println(value)
    // delete the last
    stack = stack[:len(stack)-1]
    fmt.Println(stack)
}

func main() {
    listArr()
    queueArr()
    stackArr()
}
package main

import(
	"fmt"
)


func changeParam(param int){
	param += 3
}

func changeList(param []int){
	param[0] = 99
}

func changeMap(param map[int]string){
	param[4] = "ana"
	param[1] = "adamastor"
}

func main(){
	val1 := 10
	fmt.Println("before send the var to function", val1)
	changeParam(val1)
	fmt.Println("after send the var to function", val1)
	fmt.Println("---")

	list := []int{1,2,3,4}
	fmt.Println("before send the array to function", list)
	changeList(list)
	fmt.Println("after send the array to function", list)
	fmt.Println("---")

	dict := map[int]string { 1: "joao", 2: "maria", 3: "jose",}
	fmt.Println("before send the map to function", dict)
	changeMap(dict)
	fmt.Println("after send the map to function", dict)
}
package main

import "fmt"

func functionName(){
	fmt.Println("functionName")
}

func functionWithParamsWithSameType(param1, param2 int){
	fmt.Println("functionWithParamsWithSameType", param1, param2)
}

func functionWithParamsWithSameType2(param1, param2 int, param3 string){
	fmt.Println("functionWithParamsWithSameType2", param1, param2, param3)
}

func functionReturn()int{
	fmt.Println("functionReturn")
	return 5
}

func sum(param1 int, param2 int)int{
	return param1+param2
}

func receiveNparams(list ...int)int{ // ...int transform all params in []int. declare as []int doesnt accept N params
	result := 0
	for _,v := range list{
		result += v
	}
	return result
}


func main(){
	functionName()
	functionWithParamsWithSameType(1,2)
	functionWithParamsWithSameType2(1,2,"hello")
	result := functionReturn()
	fmt.Println("value returned", result)
	result = sum(2,3)
	fmt.Println("sum value:", result)

	fmt.Println("---- testing pass N params to same function")
	fmt.Println( receiveNparams(1,2,3,4,5,6,7,8) )
	fmt.Println( receiveNparams(1,2,3) )
	fmt.Println( receiveNparams(1) )
	fmt.Println( receiveNparams() )
}
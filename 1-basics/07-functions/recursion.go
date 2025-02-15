package main

import(
	"fmt"
)


func fatorial(num int)int{
	if(num <= 1){ return num }

	return num*fatorial(num-1)
}

func main(){
	fmt.Println( fatorial(3) ) // 6
	fmt.Println( fatorial(5) ) // 120
}
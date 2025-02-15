package main

import(
	"fmt"
	"slices"
)

func functionManyReturns() (int, int){
	return 5, 8
}

func minMax(list []int) (int, int){
	return slices.Min(list), slices.Max(list)
}

func minMax2(val1, val2, val3 int) (int, int){
	list := []int {val1, val2, val3}
	return slices.Min(list), slices.Max(list)
}

// Go can named the vars to be returned in the header, so that way we dont need to create them in the body neither specify them in return
func functionWithNamedReturns(param string)(first int, last int){
	str0 := param[0]
	strLast := param[len(param)-1]

	first = int(str0)
	last = int(strLast)
	return
}

func main(){
	var1, var2 := functionManyReturns()
	fmt.Println(var1, var2)

	list := []int{4,5,2,1,8}
	var1, var2 = minMax(list)
	fmt.Println(var1, var2)

	var1, var2 = minMax2(4, 9, 2)
	fmt.Println(var1, var2)

	fmt.Println( functionWithNamedReturns("testing") )
}
package second_package

func SumArr(arr []int)int{
	total := 0
	for _, val := range(arr) {
		total += val
	}
	return total
}
package second_package

import("testing")

// test only the function in second_package file
func TestSumArr(t *testing.T){
	values := []int{3,7,6} // sum: 16
	if( SumArr(values) != 16 ){
		t.Error(values, "sum is 16")
	}

	values = []int{5,9,8} // sum: 22
	if( SumArr(values) != 22 ){
		t.Error(values, "sum is 22")
	}

	values = []int{25,31} // sum: 56
	if( SumArr(values) != 56 ){
		t.Error(values, "sum is 56")
	}
}
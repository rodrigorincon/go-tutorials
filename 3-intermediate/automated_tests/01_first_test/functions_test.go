package some_package // for unit test the test file must have in the same package
// for high level tests the package must be different, but it's very usual test private things in unit test

import("testing")

// the test file must finish with "_test"
// we must create a mod for the test comand works!
// to run the test, run "go test"
// we can use too "go test -v" to print more infor about each function ran
// OBS: run "go test FILE_TEST" fails because it doesn't find the things in the code file

// all test functions must start with "Test" and receive a single param *testing.T
// Go will list the functions that starts with "Test" to run
func TestStudentPass(t *testing.T){
	grades := []int{3,7,6} // average: 5,33333
	if( !StudentPass(grades, 5) ){
		t.Error(grades, "average must be greater or equals to 5") // we don't have a test for OK case, just for when it fails
	}

	if( StudentPass(grades, 6) ){
		t.Error(grades, "average must be less than 6")
	}

	grades = []int{2, 6}
	if( !StudentPass(grades, 4) ){
		t.Error(grades, "average must be greater or equals to 4")
	}

	grades = []int{5}
	if( !StudentPass(grades, 5) ){
		t.Error(grades, "average must be greater or equals to 5")
	}

	grades = []int{}
	if( StudentPass(grades, 0) ){
		t.Error("empty array must always fails")
	}
}

func TestIsLegalAge(t *testing.T){
	if( IsLegalAge(0) ){
		t.Error("0 is not a legal age")
	}

	if( IsLegalAge(17) ){
		t.Error("17 is not a legal age")
	}

	if( !IsLegalAge(18) ){
		t.Error("18 is a legal age")
	}

	if( !IsLegalAge(19) ){
		t.Error("19 is a legal age")
	}
}

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
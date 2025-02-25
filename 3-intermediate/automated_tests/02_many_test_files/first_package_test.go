package first_package

import("testing")

// test only the function in first_package file
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


func TestSecondTest(t *testing.T){
	grades := []int{3} 
	if( !StudentPass(grades, 3) ){
		t.Error(grades, "average must be greater or equals to 3")
	}

	if( !StudentPass(grades, 2) ){
		t.Error(grades, "average must be less than 2")
	}

	if( StudentPass(grades, 4) ){
		t.Error(grades, "average must be less than 4")
	}
}

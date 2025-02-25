package first_package

import("testing")

// test only the function in first_package file
func TestMoreStudentPass(t *testing.T){
	grades := []int{1,1,1} // average: 1
	if( !StudentPass(grades, 1) ){
		t.Error(grades, "average must be greater or equals to 1")
	}

	grades = []int{2,2} // average: 2
	if( !StudentPass(grades, 1) ){
		t.Error(grades, "average must be greater or equals to 1")
	}

	grades = []int{6} // average: 6
	if( !StudentPass(grades, 6) ){
		t.Error(grades, "average must be greater or equals to 6")
	}
}

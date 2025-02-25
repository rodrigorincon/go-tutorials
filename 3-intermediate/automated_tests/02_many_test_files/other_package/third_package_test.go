package other_package

import("testing")

// test only the function in third_package file
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

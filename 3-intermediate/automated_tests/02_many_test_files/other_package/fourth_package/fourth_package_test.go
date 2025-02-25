package fourth_package

import("testing")

// test only the function in fourth_package file
func TestFourthPkg(t *testing.T){
	if( Upper("0") != "0" ){
		t.Error("number must be the same after upper")
	}

	if( Upper("hello world") != "HELLO WORLD" ){
		t.Error("hello world was not uppercased")
	}

	if( Upper("hELLO") != "HELLO" ){
		t.Error("first letter was not uppercased")
	}
}

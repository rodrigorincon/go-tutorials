package some_package

import("testing")

// this function will not be called directly by the test, only by the test functions
func createPerson(name string, age int)Person{
	return Person{name, age}
}

// this function will be called
func TestPerson(t *testing.T){
	person := createPerson("jhon", 18)
	if( !person.IsLegalAge() ){
		t.Error(person.Age, "is legal age")
	}

	person = createPerson("mary", 5)
	if( person.IsLegalAge() ){
		t.Error(person.Age, "is not legal age")
	}

	person = createPerson("joseph", 50)
	if( !person.IsLegalAge() ){
		t.Error(person.Age, "is legal age")
	}
}
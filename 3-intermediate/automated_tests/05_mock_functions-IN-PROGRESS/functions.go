package some_package

type Person struct{
	Id int
	IdentityNumber string
	Name string
}

// function that call other function
func CreatePerson(identityNumber string, name string) *Person{
	if( !ApiValidateIdNumber(identityNumber) ){ return nil }

	return &Person{
		Id: 123,
		IdentityNumber: identityNumber,
		Name: name,
	}
}

// function to be mocked
func ApiValidateIdNumber(identity string)bool{
	// let's imagine here we have an api call to validate it
	return true
}

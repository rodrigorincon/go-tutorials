package main

import( // import searches inside GOPATH directory, so our packages must be inside there
	"fmt"
	"my_third_package" // to work we need to copy the my_third_package folder into the GOPATH directory
)

// lowercased vars, functions, structs and attributes are not exported. But it's used regurlarly inside in the functions
// to export, the first letter must be capitalized

// just for show the created package, we keep the my_third_package directory here in the same directory, but this files are not used, only the copy in GOPATH
func main(){
	fmt.Println("I'm in the main method")
	
	fmt.Println("vars imported: MyFirstVar = ", my_third_package.MyFirstVar, "MY_CONST = ", my_third_package.MY_CONST)
	my_third_package.MyMethod()

	fmt.Println("private var accessed by public method: ", my_third_package.AcessingPrivateVar())
	my_third_package.IncreasePrivateVar(3)
	fmt.Println("private var accessed by public method: ", my_third_package.AcessingPrivateVar())
	my_third_package.IncreasePrivateVar(4)
	fmt.Println("private var accessed by public method: ", my_third_package.AcessingPrivateVar())
	// my_third_package.notExportedVar // doesnt work!! var is downcased, so it's private
	// my_third_package.myPrivateFunction() // doesnt work!! function is downcased, so it's private
	
	
	var person my_third_package.MyStruct = my_third_package.MyStruct{
		Name: "joao",
		Age: 20,
		Cpf: "123.123.123-12",
		// gender: "male", we can't use gender here because it's a private attribute (downcased)
	}
	fmt.Println("person:")
	person.Print()
	// person.gender = "male" // doesnt work!! gender is downcased, so it's private
	person.SetGender("male")
	fmt.Println("person gender:", person.GetGender())
	// person.domicile = my_third_package.house{"rua aquele lugar numero 1", 40} // // doesnt work!! domicile and house are downcased, so they're private
	person.BuildHome("rua aquele lugar numero 1", 40)
	person.Print()
	fmt.Println("data is valid? ", person.Validate())
}
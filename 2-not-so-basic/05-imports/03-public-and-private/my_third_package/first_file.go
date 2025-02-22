package my_third_package

import(
	"fmt"
)

// vars and methods exported ------------------------------------------------------------
var MyFirstVar int = 5
const MY_CONST int = 200

func MyMethod(){
	fmt.Println("I'm in the exported method")
}

type MyStruct struct {
	Name string
	Age int
	Cpf string
	gender string // this attribute is not accessible in other imports
	domicile house
}
func (m *MyStruct) SetGender(gender string){
	m.gender = gender
}
func (m MyStruct) GetGender()string{
	return m.gender
}
func (m *MyStruct) BuildHome(address string, meters int){
	m.domicile = house{
		address: address,
		meters: meters,
	}
}
func (m MyStruct) Validate()bool{
	cpfValid, isAlive, gender := m.checkCpfInFederalEndpoint()
	if(!cpfValid || !isAlive || m.gender != gender){
		return false
	}
	return true
}
func (m MyStruct) Print(){
	fmt.Println("name:", m.Name, "age:", m.Age, "cpf:", m.Cpf, "gender:", m.gender, "domicile:")
	m.domicile.print()
}

func AcessingPrivateVar()int{
	return notExportedVar
}
func IncreasePrivateVar(value int){
	notExportedVar += value
	if(notExportedVar > notExportedConst){
		notExportedVar = notExportedConst
		myPrivateFunction()
	}
}

// vars and methods not exported ------------------------------------------------------------
var notExportedVar int = 11
const notExportedConst int = 15

func myPrivateFunction(){
	fmt.Println("I'm in a private function. I enter here when the notExportedVar raises the maximum value: ", notExportedConst)
}

func (m MyStruct) checkCpfInFederalEndpoint()(cpfValid bool, isAlive bool, gender string){
	return true, true, "male"
}

type house struct{
	address string
	meters int
}
func (h house) print(){
	fmt.Println("address:", h.address, "meters:", h.meters)
}
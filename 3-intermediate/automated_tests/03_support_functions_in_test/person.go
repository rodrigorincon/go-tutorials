package some_package

type Person struct{
	Name string
	Age int
}
func (p Person) IsLegalAge()bool{
	return p.Age >= 18
}
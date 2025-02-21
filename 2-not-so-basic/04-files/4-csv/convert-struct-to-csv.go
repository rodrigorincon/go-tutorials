package main

import(
	"os"
	"bufio"
	"strings"
	"strconv"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

type Person struct{
	Name string
	Surname string
	Cellphone string
	City string
	State string
	Age int
	UseGlasses bool
}

func main(){
	// define the objects
	joao := Person{Name: "João", Surname: "Silva", Cellphone: "85987654321", City: "Austin", State: "TX", Age: 30, UseGlasses: false}
	mario := Person{Name: "Mario", Surname: "Andrade", Cellphone: "6912345678", City: "New Youk", State: "NY", Age: 10, UseGlasses: true}
	ana := Person{Name: "Ana", Surname: "Sousa", Cellphone: "6189674532", City: "Los Angeles", State: "CA", Age: 20, UseGlasses: true}
	jose := Person{Name: "José", Surname: "Oliveira", Cellphone: "5585973612", City: "Boston", State: "MA", Age: 6, UseGlasses: false}
	maria := Person{Name: "Maria", Surname: "Silveira", Cellphone: "729384756", City: "Houston", State: "TX", Age: 28, UseGlasses: false}
	people := []Person{joao, mario, ana, jose, maria}

	// create the file
	file, err := os.Create("people-sheet.csv")
    checkError(err)
	writer := bufio.NewWriter(file)

	// define csv attributs
	columnDelim := ","
	lineDelim := "\n"
	header := []string{"name", "surname", "cellphone", "city", "state", "age", "use glasses"}
	
	// write the header
	_, err = writer.WriteString( strings.Join(header, columnDelim) + lineDelim )
	checkError(err)

	// add the content
	for _, person := range(people) {
		hasGlasses := "No"
		if(person.UseGlasses){
			hasGlasses = "Yes"
		}
		personArr := []string{ person.Name, person.Surname, person.Cellphone, person.City, person.State, strconv.Itoa(person.Age), hasGlasses }
		_, err = writer.WriteString( strings.Join(personArr, columnDelim) + lineDelim )
		checkError(err)
	}

	writer.Flush()
	file.Close()
}
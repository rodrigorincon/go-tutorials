package main

import(
	"fmt"
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
	IsLegalAge bool
}

func main(){
	// open the file
	file, err := os.Open("csv-to-struct.csv")
    checkError(err)
	scanner := bufio.NewScanner(file)

	// define csv attributs
	hasHeader := true
	columnDelim := ","
		
	// skip the header
	if(hasHeader){
		scanner.Scan()	
	}
	
	// read the content
	for scanner.Scan() {
		line := scanner.Text()

		attributes := strings.Split(line, columnDelim)

		age, err2 := strconv.Atoi(attributes[5])
	    checkError(err2)

		hasGlasses := false
		if(attributes[6] == "Yes"){
			hasGlasses = true
		}
		
		person := Person{ 
			Name: attributes[0],
			Surname: attributes[1],
			Cellphone: attributes[2],
			City: attributes[3],
			State: attributes[4],
			Age: age,
			UseGlasses: hasGlasses,
			IsLegalAge: age >= 18,
		}
		fmt.Println(person)
	}

	file.Close()
}
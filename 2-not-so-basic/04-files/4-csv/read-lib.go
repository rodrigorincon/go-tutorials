package main

import(
	"fmt"
	"os"
	"io"
	"encoding/csv"
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
	
	// create csv reader object
	reader := csv.NewReader(file)
	reader.Comma = ','
	hasHeader := true
		
	// skip the header
	if(hasHeader){
		_, err = reader.Read()	
		checkError(err)
	}
	
	// read the content
	finishRead := false
	for !finishRead {
		attributes, err2 := reader.Read()
		if( err2 == io.EOF){
			finishRead = true
		}else{
			checkError(err2)
		}

		if( len(attributes) == 0 ){
			break
		}

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
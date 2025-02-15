package main

import(
	"fmt"
)

func filterPart(){
	fmt.Println("--------------- Filter function")
	filterArray := func(list []int, filterFunc func(int)bool)(resultList []int){
		for _, value := range list{
			if( filterFunc(value) ){
				resultList = append(resultList, value)
			}
		}
		return
	}
	myArray := []int{1,2,3,4,5,6,7,8,9,0}
	evenNumbers := filterArray(myArray, func(value int)bool{
		return value % 2 == 0 // return only even numbers
	})
	fmt.Println("original:", myArray)
	fmt.Println("filtered:", evenNumbers)
}

func mapPart(){
	fmt.Println("--------------- Map function")
	mapArray := func(list []string, mapFunc func(string)int)(resultList []int){
		for _, value := range list{
			mappedValue := mapFunc(value)
			resultList = append(resultList, mappedValue)
		}
		return
	}
	myArray2 := []string{"a","b","c","d"}
	ascIIChars := mapArray(myArray2, func(letter string)int{
		return int([]byte(letter)[0]) // return the ascii number of the letter
	})
	fmt.Println("original:", myArray2)
	fmt.Println("filtered:", ascIIChars)
}

func mapStructPart(){
	fmt.Println("--------------- Map for struct function")
	type Person struct{
		name string
		age int
	}
	mapStruct := func(list []Person, mapFunc func(Person)int)(resultList []int){
		for _, value := range list{
			mappedValue := mapFunc(value)
			resultList = append(resultList, mappedValue)
		}
		return
	}
	p1 := Person{"joao", 12}
	p2 := Person{"jose", 18}
	p3 := Person{"maria", 7}
	p4 := Person{"ana", 20}
	people := []Person{p1,p2,p3,p4}
	ages := mapStruct(people, func(person Person)int{
		return person.age // return the age of each person
	})
	fmt.Println("original:", people)
	fmt.Println("filtered:", ages)
}

func main(){
	// implementing a filter function
	filterPart()

	// implementing a map function
	mapPart()

	// implementing other map function
	mapStructPart()
}
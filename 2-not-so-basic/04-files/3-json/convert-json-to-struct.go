package main

import(
	"fmt"
	"os"
	"encoding/json"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

type Person struct{
	Name string
	Age int
	UseGlasses bool
	Friends []Person
}

func main(){
	// read at once
	fileBytes, err := os.ReadFile("json-file.json") // this method opens, read the complete file and closes. The data returned is in bytes
    checkError(err)
    
    newPeople := []Person{}
    json.Unmarshal( fileBytes, &newPeople) // it worked independent of the tags or if the keys are downcased or capital lettered
	fmt.Println("how many people was created:", len(newPeople))
	for _, person := range(newPeople) {
		fmt.Println("name:", person.Name, "age:", person.Age, "use glasses:", person.UseGlasses)
		for _, friend := range(person.Friends){
			fmt.Println("name:", friend.Name, "age:", friend.Age, "use glasses:", friend.UseGlasses)
		}
		fmt.Println("------")
	}
}
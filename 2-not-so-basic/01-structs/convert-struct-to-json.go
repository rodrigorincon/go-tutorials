package main

import(
	"fmt"
	"encoding/json"
	"unsafe"
)

type Person struct{
	Name string
	Age int
	Single bool
}
// without the tag the json is {"Name":"joao","Age":10,"Single":true} with keys capital lettered

type PersonTagged struct{
	Name string `json:"name"` // this `json:` is a tag. We add tags in attributes to rename the attribute in specific formats
	Age int     `json:"age"`
	Single bool `json:"single" xml:"single"`
}
// with the tag the json is {"name":"joao","age":10,"single":true} with keys lowercased

func main(){
	person := Person{Name: "joao", Age: 10, Single: true}
	fmt.Println("person: ", person)
	fmt.Println("untagged size: ", unsafe.Sizeof(person) )

	// convert to json
	jsonPerson, err := json.Marshal(person)
	if(err != nil){
		fmt.Println("error: ", err.Error())
		return
	}
	// json.Marshal return an array of bytes
	fmt.Println( jsonPerson )
	// we need to convert to string to see
	fmt.Println( string(jsonPerson) )

	fmt.Println("--------- Tagged struct")

	personTag := PersonTagged{Name: "joao", Age: 10, Single: true}
	fmt.Println("person Tag: ", personTag)
	fmt.Println("tagged size: ", unsafe.Sizeof(personTag) ) // doesnt change the size

	jsonPersonTag, errTag := json.Marshal(personTag)
	if(errTag != nil){
		fmt.Println("error: ", errTag.Error())
		return
	}
	// we need to convert to string to see
	fmt.Println( string(jsonPersonTag) )


	fmt.Println("--------- build a struct by a json")

	rawJson := `{"name":"joao","age":10,"single":true}`
	newPerson := Person{}

	json.Unmarshal( []byte(rawJson), &newPerson) // it worked independent of the tags or if the keys are downcased or capital lettered
	fmt.Println("my new person:", newPerson.Name, newPerson.Age, newPerson.Single)

}
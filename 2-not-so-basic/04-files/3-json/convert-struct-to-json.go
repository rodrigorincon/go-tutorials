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
	Name string			`json:"name"`
	Age int 			`json:"age"`
	UseGlasses bool		`json:"useglasses"`
	Friends []Person	`json:"friends"`
}

func main(){
	mario := Person{Name: "Mario de Andrade", Age: 10, UseGlasses: true}
	marioFriend1 := Person{Name: "Amigo do Mario", Age: 9, UseGlasses: false}
	marioFriend2 := Person{Name: "Amigo2 do Mario", Age: 8, UseGlasses: true}
	mario.Friends = []Person{marioFriend1, marioFriend2}

	joana := Person{Name: "Joana Oliveira", Age: 30, UseGlasses: false}
	joanaFriend1 := Person{Name: "Amigo da joana", Age: 29, UseGlasses: false}
	joanaFriend2 := Person{Name: "Amigo2 da joana", Age: 28, UseGlasses: true}
	joana.Friends = []Person{joanaFriend1, joanaFriend2}

	// json.Marshal return an array of bytes
	jsonPeopleBytes, err := json.Marshal( []Person{mario, joana} )
	if(err != nil){
		fmt.Println("error: ", err.Error())
		return
	}

	file, err2 := os.Create("new-json-file.json")
    if(err2 != nil){
		fmt.Println("error: ", err2.Error())
		return
	}

	numBytes, err3 := file.Write(jsonPeopleBytes)
	if(err3 != nil){
		fmt.Println("error: ", err3.Error())
		return
	}
	file.Close()

	fmt.Println("new file has ", numBytes, "bytes")
}
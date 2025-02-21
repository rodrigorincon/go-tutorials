package main

import(
	"fmt"
	"os"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	// creates a new file
	err := os.Rename("file-to-be-renamed.txt", "new-name.txt")
    if err != nil {
		fmt.Println("error:", err)
		return
    }
    fmt.Println("file renamed!")
}
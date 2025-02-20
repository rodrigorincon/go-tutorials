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
	// read at once
	fileBytes, err := os.ReadFile("file.txt") // this method opens, read the complete file and closes. The data returned is in bytes
    checkError(err)
    fmt.Println(fileBytes) // print array of bytes
    fmt.Println( string(fileBytes) ) // print the text
	// fileBytes.Close() // it's not necessary close the object

	// read from an already opened file
	file, err2 := os.Open("file.txt")
    checkError(err2)
    fi, err3 := file.Stat()
	checkError(err3)

    byteList := make([]byte, fi.Size())
    numBytes, err4 := file.Read(byteList)
    checkError(err4)
    fmt.Println("number of bytes:", numBytes, "number of size in FI:", fi.Size(), "text:", string(byteList) )
    file.Close()
}
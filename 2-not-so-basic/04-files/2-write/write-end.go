package main

import(
	"fmt"
    "os"
    "bufio"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
    // it's not possible define a place to write in a file without override all the content (expect the end)
    newText := "\nthis is the new text to be added. HELLO WORLD!"

    // open the file
    file, err := os.OpenFile("write-end.txt", os.O_WRONLY | os.O_APPEND, os.ModeAppend) // open to write. By default, open is read-only
    // O_WRONLY give permission to write, O_APPEND says to not override, but append the new text in the end (without it write in the being overriding)
    checkError(err)
    
    // set writer object
    writer := bufio.NewWriter(file)

    // write the new text
    numCharWriten, err2 := writer.Write( []byte(newText) )
    checkError(err2)
    fmt.Println("writen ", numCharWriten, "of", len(newText), "chars")
    
    // ensure the writing and close the file
    err3 := writer.Flush()
    checkError(err3)
    file.Close()
}
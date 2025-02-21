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
    // so we need to use 2 files and copy the original content and add the new content where we want
    newText := "this is the new text to be added. HELLO WORLD!\n"

    // open the file
    file, err := os.OpenFile("write-begin.txt", os.O_RDWR, os.ModeAppend) // open to read and write. By default, open is read-only
    // O_RDWR gives permission to read and write
    checkError(err)

    // read the content and store the whole file content in the memory (not good)
    originalText := ""
    reader := bufio.NewScanner(file)
    for reader.Scan() {
        originalText += reader.Text()+"\n"
    }
    originalText = originalText[:len(originalText)-1] // remove the last "\n" added in the loop
    fmt.Println("original file:\n", originalText)

    // build the final text
    finalText := newText + originalText
    
    // set writer object
    file.Seek(0, 0) // rewind
    writer := bufio.NewWriter(file)

    // write the new text
    numCharWriten, err2 := writer.Write( []byte(finalText) )
    checkError(err2)
    fmt.Println("writen ", numCharWriten, "of", len(finalText), "chars")
    
    // ensure the writing and close the file
    err3 := writer.Flush()
    checkError(err3)
    file.Close()
}
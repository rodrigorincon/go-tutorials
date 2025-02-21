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
	// write in a file. If file doesn't exist, create it. If exist, overwrite the content
	newContent := []byte("hello world\n")
    err := os.WriteFile("new-file.txt", newContent, 0644) // last param is the permission (rw-r-r)
    checkError(err)

    err = os.WriteFile("already-existed-file.txt", newContent, 0644)
    checkError(err)


    file, err := os.Create("new-file2.txt") // create a new file. If already exist, open the existed
    checkError(err)

    newContent = []byte("hello world\n")
    numBytes, err2 := file.Write(newContent) // if file already exists, overwrite the content
    checkError(err2)
    fmt.Println("bytes wroted:", numBytes)

    numBytes2, err3 := file.WriteString("a new text")
    checkError(err3)
    fmt.Println("bytes wroted:", numBytes2)

    file.Sync() // flush the file, enure the internal buffers will be released and file data will be stored
    file.Close()
}
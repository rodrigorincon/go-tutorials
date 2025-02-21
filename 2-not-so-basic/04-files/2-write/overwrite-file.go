package main

import(
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
}
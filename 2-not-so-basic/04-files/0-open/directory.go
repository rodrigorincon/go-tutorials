package main

import(
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	fullPathFile, err := os.Executable()
	checkError(err)

	// read the complete path of the file
	fmt.Println(fullPathFile)

	// read the complete path of the folder
	dir := filepath.Dir(fullPathFile)
	fmt.Println(dir)

	// directory name
	pathArr := strings.Split(dir, "/")
	fmt.Println("folder name:", pathArr[ len(pathArr)-1 ]  )

	// concatenate folders
	fullPath := filepath.Join(dir, "..", "1-read", "log.txt")
	fmt.Println("new path:", fullPath)
}
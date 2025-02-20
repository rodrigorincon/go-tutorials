package main

import(
	"fmt"
	"os"
	"time"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	bigFile, err := os.Open("../1-read/bigfile.txt")
	checkError(err)
	
	fi, err2 := bigFile.Stat() // returns os.FileInfo object
	checkError(err2)

	var name string = fi.Name() // just the name, not the entire path
	var size int64 = fi.Size() // size in bytes
	var permitions os.FileMode = fi.Mode() // -rw-rw-r--
	var lastModificTime time.Time = fi.ModTime()
	var isDir bool = fi.IsDir()

	fmt.Println("name:", name, "size in bytes:", size, "permitions:", permitions, "last modification", lastModificTime, "is a dir?", isDir)

	err3 := bigFile.Close()
	checkError(err3)
}
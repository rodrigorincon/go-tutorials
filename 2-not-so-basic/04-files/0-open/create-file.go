package main

import(
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	// creates a new file
	file, err := os.Create("newfoo.txt") // file is a pointer
	checkError(err)
	fmt.Println("file:", file, "type of file:", reflect.TypeOf(file), "pointer size:", unsafe.Sizeof(file), "file size:", unsafe.Sizeof(*file)) // file: &{0xc0000321e0} type of file: *os.File pointer size: 8 file size: 8
	
	// close the file
	err2 := file.Close()
	checkError(err2)
}
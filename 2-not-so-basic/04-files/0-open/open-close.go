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
	// open a file that exists
	file, err := os.Open("file.txt") // file is a pointer
	checkError(err)
	fmt.Println("file:", file, "type of file:", reflect.TypeOf(file), "pointer size:", unsafe.Sizeof(file), "file size:", unsafe.Sizeof(*file)) // file: &{0xc0000321e0} type of file: *os.File pointer size: 8 file size: 8
	
	// close the file
	err2 := file.Close()
	checkError(err2)
	file.Close() // execute close for an already closed file doesnt do anything

	// open a file that doesnt exists
	file2, err3 := os.Open("file-doesnt-exist.txt")
	fmt.Println("error message: ", err3, "\nFile2 is null?", file2 ==nil)

	// open a file in other folder
	bigFile, err4 := os.Open("../1-read/bigfile.txt")
	checkError(err4)
	fmt.Println("pointer size:", unsafe.Sizeof(bigFile), "file size:", unsafe.Sizeof(*bigFile)) // size is always 8bytes (because the machines is 64bits)
	// close the file
	bigFile.Close()
}
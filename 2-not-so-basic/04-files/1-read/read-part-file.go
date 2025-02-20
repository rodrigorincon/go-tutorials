package main

import(
	"fmt"
	"os"
	"io"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	file, err := os.Open("bigfile.txt")
    checkError(err)

	// read only 10 bytes
    byteList := make([]byte, 10)
    numBytes, err2 := file.Read(byteList)
    checkError(err2)
    fmt.Println("number of bytes:", numBytes, "text:", string(byteList) ) // show the first 10 bytes

    // read the next 10 bytes
    numBytes, err2 = file.Read(byteList)
    checkError(err2)
    fmt.Println("number of bytes:", numBytes, "text:", string(byteList) ) // show from 10th to 20th byte

    // read 10 bytes after 6 bytes from the beginning of file
    index, err3 := file.Seek(6, io.SeekStart) // set the cursor to read from a position and return it. In this case, io.SeekStart (0) + 6
    checkError(err3)
    numBytes2, err4 := file.Read(byteList)
    checkError(err4)
    fmt.Println("reading from position", index, "to ", numBytes2, ":", string(byteList) )
    // we can use io.SeekCurrent to get the current position and io.SeekEnd to get the last position (with a negative number to read the last ones)

    file.Close()
}
package main

import(
	"fmt"
    "os"
    "io"
    "bufio"
)

func checkError(e error){
	if(e != nil && e != io.EOF){
		panic(e)
	}
}

const BUFFER_SIZE int = 40

func main(){
    // it's not possible define a place to write in a file without override all the content (expect the end)
    // so we need to use 2 files and copy the original content and add the new content where we want
    newText := "this is the new text to be added. HELLO WORLD!\n"

    // open the file to read and their vars
    fileRead, err := os.Open("write-begin.txt")
    checkError(err)
    fi, err2 := fileRead.Stat()
    checkError(err2)
    byteSize := fi.Size()
    reader := bufio.NewReaderSize(fileRead, BUFFER_SIZE)

    // open the file to write and their vars
    fileWrite, err3 := os.Create("new-write-begin.txt")
    checkError(err3)
    writer := bufio.NewWriterSize(fileWrite, BUFFER_SIZE)

    // write the new text in the write file in batches
    for index := 0; index < len(newText); index += BUFFER_SIZE {

        maxIndex := index+BUFFER_SIZE
        if( maxIndex > len(newText) ){
            maxIndex = len(newText)
        }

        substring := newText[index : maxIndex]
        _, err4 := writer.Write( []byte(substring) )
        checkError(err4)

        fmt.Println(substring, "\n")
    }

    // write the content of the current file in the new file in batches
    buffer := make([]byte, BUFFER_SIZE)
    for index := 0; int64(index) < byteSize; {
        numBytesRead, err4 := reader.Read(buffer)
        checkError(err4)
        if(err4 == io.EOF){ break }

        _, err4 = writer.Write(buffer[:numBytesRead])
        checkError(err4)

        index += numBytesRead

        fmt.Println(string(buffer[:numBytesRead]), "\n")
    }
    
    // ensure the writing and close the files
    err4 := writer.Flush()
    checkError(err4)
    fileWrite.Close()
    fileRead.Close()

    err5 := os.Remove("write-begin.txt")
    checkError(err5)
    
    os.Rename("new-write-begin.txt", "write-begin.txt")
}
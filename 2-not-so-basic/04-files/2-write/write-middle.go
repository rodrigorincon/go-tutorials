package main

import(
	"fmt"
    "os"
    "bufio"
    "io"
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
    newText := " THIS IS THE NEW TEXT TO BE ADDED. HELLO WORLD "
    indexToInsert := 135

    // open the file to read and their vars
    fileRead, err := os.Open("write-middle.txt")
    checkError(err)
    fi, err2 := fileRead.Stat()
    checkError(err2)
    byteSize := fi.Size() + int64(len(newText))
    reader := bufio.NewReaderSize(fileRead, BUFFER_SIZE)

    // open the file to write and their vars
    fileWrite, err3 := os.Create("new-write-middle.txt")
    checkError(err3)
    writer := bufio.NewWriterSize(fileWrite, BUFFER_SIZE)

    // write the content of the current file in the new file AND in the new content in batches
    buffer := make([]byte, BUFFER_SIZE)
    for index := 0; int64(index) < byteSize; {
        numBytesRead, err4 := reader.Read(buffer)
        checkError(err4)
        if(err4 == io.EOF){ break }

        usedBuffer := buffer[:numBytesRead]
        if( indexToInsert >= index && indexToInsert <= index+numBytesRead ){
            bufferIndex := indexToInsert - index

            firstPart := usedBuffer[:bufferIndex]
            _, err4 = writer.Write(firstPart)
            checkError(err4)
            fmt.Println(string(firstPart), "\n")

            // write the new text in the write file in batches
            for indexNew := 0; indexNew < len(newText); indexNew += BUFFER_SIZE {

                maxIndex := indexNew+BUFFER_SIZE
                if( maxIndex > len(newText) ){
                    maxIndex = len(newText)
                }

                substring := newText[indexNew : maxIndex]
                _, err4 = writer.Write( []byte(substring) )
                checkError(err4)

                fmt.Println(substring, "\n")
            }

            lastPart := usedBuffer[bufferIndex:]
            _, err4 = writer.Write(lastPart)
            checkError(err4)
            fmt.Println(string(lastPart), "\n")
        }else{
            _, err4 = writer.Write(usedBuffer)
            checkError(err4)
            fmt.Println(string(usedBuffer), "\n")
        }
        index += numBytesRead
    }
    
    // ensure the writing and close the files
    err4 := writer.Flush()
    checkError(err4)
    fileWrite.Close()
    fileRead.Close()

    err5 := os.Remove("write-middle.txt")
    checkError(err5)
    
    os.Rename("new-write-middle.txt", "write-middle.txt")
}
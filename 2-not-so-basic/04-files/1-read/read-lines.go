package main

import(
	"fmt"
	"os"
    "bufio"
    "io"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	file, err := os.Open("file-lines.txt")
    checkError(err)

    // FIRST WAY - USING SCANNER
	scanner := bufio.NewScanner(file) // create a buffer with some useful methods to read
    // scanner.Buffer(mySlice, size) // set a buffer size to be used in Scan(). The default size is 64k

    hasMore := scanner.Scan() // read until find a defined byte. By default is "\n"
    for hasMore {
        fmt.Println("after scan:", scanner.Text(), "read everyhing?", hasMore)
        // we can get the bytes instead the string using scanner.Bytes()
        hasMore = scanner.Scan()
    }

    file.Seek(0, io.SeekStart)
    // Reader ----------------------------
    fmt.Println("-----------")

    // SECOND WAY - USING READER
    reader := bufio.NewReader(file) // create another buffer with other useful methods to read
    // reader.NewReaderSize(file, size) // define a size for the buffer

    fileEnded := false
    for !fileEnded {
        line, hasMoreContentLine, err2 := reader.ReadLine() // read the line, but if the buffer is not big enough, the second response is true
        // we can use Reader.ReadBytes('\n') or Reader.ReadString('\n') as the same way
        if err2 == io.EOF {
            fileEnded = true
            break
        }else if err2 != nil {
            panic(err2)
        }

        for hasMoreContentLine { // if the line was not completely read
            var restLine []byte
            restLine, hasMoreContentLine, err2 = reader.ReadLine() // continue reading

            if err2 == io.EOF {
                fileEnded = true
            }else if err2 != nil {
                panic(err2)
            }
            
            line = append(line, restLine...) // add the new part in the line
        }
        fmt.Println("after scan:", string(line))
    }
    
    file.Close()
}
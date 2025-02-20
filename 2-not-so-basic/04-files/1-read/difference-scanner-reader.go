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

// use SCANNER when you know the size of the lines and the lines have same sized lines
// use READER when the file is inconstant sized
func main(){
	file, err := os.Open("file-lines.txt")
    checkError(err)

    // SCANNER uses a predefined buffer size. It's useful when you know nearly the size of each line and when all lines have nearly the same size
	scanner := bufio.NewScanner(file) 
    // scanner.Buffer(mySlice, size) 

    hasMore := scanner.Scan() 
    for hasMore {
        fmt.Println(scanner.Text()) // scanner.Bytes()
        hasMore = scanner.Scan()
    }

    file.Seek(0, io.SeekStart)
    // Reader ----------------------------
    fmt.Println("-----------")

    // READER doesnt use a predefined buffer size. The buffer size varies in each loop according the necessity. Spends less memory
    reader := bufio.NewReader(file)
    // reader.NewReaderSize(file, size) // define a size for the buffer

    fileEnded := false
    for !fileEnded {
        line, err2 := reader.ReadString('\n') // Reader.ReadBytes('\n')
        if err2 == io.EOF {
            fileEnded = true
        }else if err2 != nil {
            panic(err2)
        }

        fmt.Println(line)
    }
    
    file.Close()
}
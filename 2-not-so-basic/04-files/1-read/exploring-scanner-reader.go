package main

import(
	"fmt"
	"os"
    "bufio"
    "io"
    "bytes"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func readUntilLetterI(data []byte, atEOF bool) (advance int, token []byte, err error){
    if(atEOF && len(data) == 0) {
        return 0, nil, nil
    }
    if index := bytes.IndexByte(data, 'i'); index >= 0 {
        return index + 1, data[0:index], nil
    }
    // If we're at EOF, we have a final, non-terminated line. Return it.
    if atEOF {
        return len(data), data, nil
    }
    // Request more data.
    return 0, nil, nil
}


func scannerFunc(file *os.File){
    fmt.Println("------------ reading with SCANNER until end of line")
    scanner := bufio.NewScanner(file) // create a buffer with some useful methods to read

    myBuffer := make([]byte, 50) // as my file doesnt have so much letters in a line, we can define a small buffer. 
    // If necessary we can define a bigger buffer than the default too
    
    scanner.Buffer(myBuffer, len(myBuffer)) // set a buffer size to be used in Scan(). The default size is 64k

    for scanner.Scan() { // read until find a defined byte. By default is "\n"
        fmt.Println(scanner.Text())
    }
    // back to the beginning of file
    file.Seek(0, io.SeekStart)
    
    fmt.Println("\n\n------------ reading with SCANNER until letter i")

    scanner = bufio.NewScanner(file) // we can't rewind the scanner, we need to create a new scanner

    // change how far the Scan() method read, passing the function that define the char used to stop the reading
    scanner.Split( readUntilLetterI )
    for scanner.Scan() { // read until find a defined byte. By default is "\n"
        fmt.Println(scanner.Text())
        fmt.Println("\n")
    }
    // back to the beginning of file
    file.Seek(0, io.SeekStart)

    fmt.Println("\n\n------------ reading WORDS with SCANNER")

    scanner = bufio.NewScanner(file) // we can't rewind the scanner, we need to create a new scanner

    // change how far the Scan() method read, passing the function that define the char used to stop the reading
    scanner.Split( bufio.ScanWords )
    for scanner.Scan() { // read until find a defined byte. By default is "\n"
        fmt.Println(scanner.Text())
    }
    // back to the beginning of file
    file.Seek(0, io.SeekStart)
}

func readerFunc(file *os.File){
    fmt.Println("------------ reading with READER until end of line")
    reader := bufio.NewReader(file)

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
    // back to the beginning of file
    file.Seek(0, io.SeekStart)
    fmt.Println("------------ reading with READER until letter i")
    
    reader = bufio.NewReader(file) // we can't rewind the reader, we need to create a new reader
    fileEnded = false
    for !fileEnded {
        line, err2 := reader.ReadString('i') // Reader.ReadBytes('i')
        if err2 == io.EOF {
            fileEnded = true
        }else if err2 != nil {
            panic(err2)
        }
        fmt.Println(line)
    }
    // back to the beginning of file
    file.Seek(0, io.SeekStart)
    fmt.Println("\n\n------------ reading WORDS with READER")

    reader = bufio.NewReader(file) // we can't rewind the reader, we need to create a new reader
    fileEnded = false
    for !fileEnded {
        word, err2 := reader.ReadString(' ') // Reader.ReadBytes(' ')
        if err2 == io.EOF {
            fileEnded = true
        }else if err2 != nil {
            panic(err2)
        }
        fmt.Println(word)
    }
    // back to the beginning of file
    file.Seek(0, io.SeekStart)
}

func main(){
	file, err := os.Open("file-lines.txt")
    checkError(err)
    scannerFunc(file)
    readerFunc(file)

    file.Close()
}
package main

import(
	"fmt"
	"os"
    "io"
    "bufio"
    "strings"
)

func checkError(e error){
	if(e != nil){
		panic(e)
	}
}

func main(){
	file, err := os.Open("log.txt")
    checkError(err)
    
    fmt.Println("------------ searching with SCANNER")
    scanner := bufio.NewScanner(file)

    for scanner.Scan() { // read line
        line := scanner.Text()
        contains := strings.Contains(line, "/other/4321/path/method")
        if(contains){
            fmt.Println("The line with the searched part:", line)
            break
        }
    }

    // back to the beginning of file
    file.Seek(0, io.SeekStart)
    fmt.Println("------------ searching with READER")

    reader := bufio.NewReader(file)
    for {
        line, err2 := reader.ReadString('\n') 
        if err2 == io.EOF {
            break
        }else if err2 != nil {
            panic(err2)
        }
        
        contains := strings.Contains(line, "/other/4321/path/method")
        if(contains){
            fmt.Println("The line with the searched part:", line)
            break
        }
    }
 

    file.Close()
}
package main

import(
	"fmt"
	"encoding/binary"
	"reflect"
)

func main(){
	var number int32 = 259
	var text string = "Hi!" // 72 105 33 

	// get the bytes of the number
	bytesQtt := reflect.TypeOf(number).Size() // get number of bytes of the number
	
	numberBytes := make([]byte, bytesQtt)
    binary.BigEndian.PutUint32(numberBytes, uint32(number))
    fmt.Println(numberBytes)

	// get the bytes of the string
	textBytes := []byte(text)
	fmt.Println(textBytes)
}
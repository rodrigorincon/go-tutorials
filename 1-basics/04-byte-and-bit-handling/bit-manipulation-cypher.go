package main

import(
	"fmt"
)

func main(){
	var text string = "Hello world 😊!"
	// get the bytes of the string
	textBytes := []byte(text)
	fmt.Println(textBytes)

	// cypher rotating the bit to left and the first to last position
	for index := 0; index < len(textBytes); index++ {
		myByte := textBytes[index]
		onlyFirstBit := myByte & 128 // AND operation with 10000000. This line returns or 10000000 or 00000000
		firstBit := onlyFirstBit >> 7 // returns 1 or 0

		newValue := (myByte << 1) + firstBit
		textBytes[index] = newValue
	}
	fmt.Println(textBytes)
	fmt.Println( string(textBytes) )

	// decypher back
	for index := 0; index < len(textBytes); index++ {
		myByte := textBytes[index]
		onlyLastBit := myByte & 1 // AND operation with 00000001. This line returns or 00000001 or 00000000
		lastBit := onlyLastBit << 7 // returns 128 or 0

		newValue := (myByte >> 1) + lastBit
		textBytes[index] = newValue
	}
	fmt.Println(textBytes)
	fmt.Println( string(textBytes) )
}
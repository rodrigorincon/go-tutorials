package main

import(
	"fmt"
)

func main(){
	var text string = "Hello world 😊!"
	// get the bytes of the string
	textBytes := []byte(text)
	fmt.Println(textBytes)

	// cypher adding 3 in each byte. If the byte pass 255 the rest is the value (ex: 254+3 = 2)
	for index := 0; index < len(textBytes); index++ {
		newValue := textBytes[index] + 3 
		if newValue > 255 {
			newValue -= 255
		}
		textBytes[index] = newValue
	}
	fmt.Println(textBytes)
	fmt.Println( string(textBytes) )

	// decypher back
	for index := 0; index < len(textBytes); index++ {
		newValue := textBytes[index] - 3
		if newValue < 0 {
			newValue += 255
		}
		textBytes[index] = newValue
	}
	fmt.Println(textBytes)
	fmt.Println( string(textBytes) )
}
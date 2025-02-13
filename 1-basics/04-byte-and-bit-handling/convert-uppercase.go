package main

import(
	"fmt"
)

func main(){
	var text string = "Hello World 123!"
	// get the bytes of the string
	textBytes := []byte(text)
	fmt.Println(textBytes)

	// converting upper and case and the oposite
	for index := 0; index < len(textBytes); index++ {
		if textBytes[index] > 64 && textBytes[index] < 91 { //converte de maiusculo pra minusculo
			textBytes[index] += 32;
		} else if textBytes[index] > 96 && textBytes[index] < 123 { //converte de minusculo para maisuculo
			textBytes[index] -= 32;
		}
	}
	fmt.Println(textBytes)
	fmt.Println( string(textBytes) )
}
package main

import(
	"fmt"
	"strconv"
)

func hexaDigit(digit uint8)string{
	if digit < 10 { return strconv.Itoa(int(digit)) }
	switch digit {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return "0"
	}
}

func convertHexa(number uint8)string{
	var hexaNumber string = ""
	for number > 16 {
		division := uint8(number/16)		

		rest := number % 16
		hexaRest := hexaDigit(rest)
		
		hexaNumber = hexaRest+hexaNumber
		number = division
	}
	hexaRest := hexaDigit(number)
	hexaNumber = hexaRest+hexaNumber
	return hexaNumber
}

func convertOcta(number uint8)string{
	var octaNumber string = ""
	for number > 8 {
		division := uint8(number/8)		
		rest := strconv.Itoa(int(number % 8))

		octaNumber = rest+octaNumber
		number = division
	}
	octaNumber = strconv.Itoa(int(number))+octaNumber
	return octaNumber
}

func main(){
	fmt.Println("-------- CONVERTING TO DIFFERENT BASES")
	var number uint8 = 140
	// built-in way
	fmt.Println("Hexadecimal", strconv.FormatInt(int64(number), 16) )
	fmt.Println("Octadecimal", strconv.FormatInt(int64(number), 8) )
	fmt.Println("Binary", strconv.FormatInt(int64(number), 2) )
	// my way
	fmt.Println("Hexadecimal", convertHexa(number) )
	fmt.Println("Octadecimal", convertOcta(number) )
}
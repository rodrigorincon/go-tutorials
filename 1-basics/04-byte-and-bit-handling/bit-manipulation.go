package main

import(
	"fmt"
	"strconv"
)

func printBits(number uint8) []uint8 {
	bit1 := (number & 128)>>7 // first bit
	bit2 := (number & 64)>>6 // second bit
	bit3 := (number & 32)>>5 // third bit
	bit4 := (number & 16)>>4 // fourth bit
	bit5 := (number & 8)>>3 // fifth bit
	bit6 := (number & 4)>>2 // sixth bit
	bit7 := (number & 2)>>1 // seventh bit
	bit8 := (number & 1) // eight bit
	return []uint8{bit1,bit2,bit3,bit4,bit5,bit6,bit7,bit8}
}

func main(){
	fmt.Println("-------- PRINTING BITS")
	var number1 uint8 = 12 // 0000 1100
	var number2 uint8 = 54 // 0011 0110
	// built-in way
	fmt.Println(strconv.FormatInt(int64(number1), 2))
	fmt.Println(strconv.FormatInt(int64(number2), 2))
	// my way
	fmt.Println( printBits(number1) )
	fmt.Println( printBits(number2) )
	
	var resultNumber uint8
	fmt.Println("-------- BITWISE OPERATION")

	resultNumber = number1 & number2
	fmt.Println("AND ", resultNumber, printBits(resultNumber))
	resultNumber = number1 | number2
	fmt.Println("OR ", resultNumber, printBits(resultNumber))
	resultNumber = number1 ^ number2
	fmt.Println("XOR ", resultNumber, printBits(resultNumber))
	resultNumber = ^number1
	fmt.Println("NOT in ", number1, resultNumber, printBits(resultNumber))
}
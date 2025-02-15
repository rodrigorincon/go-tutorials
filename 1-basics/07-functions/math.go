package main

import(
	"fmt"
	"time"
	"math"
	"math/rand"
)


func main(){
	fmt.Println( math.Pi )
	fmt.Println( math.Pow(2,3) )
	fmt.Println( math.Sqrt(9) )
	fmt.Println( math.Abs(-4), math.Abs(5) )

	// random values
	rand.Seed(time.Now().UnixNano()) 

	fmt.Println("--- random values")
	fmt.Println(rand.Intn(10)) // the numbers are random even dont start a seed, but it's always better start it
	fmt.Println(rand.Intn(10)) // 0 to 9
	fmt.Println(rand.Intn(10))
	fmt.Println("--- rolling a dice")
	fmt.Println(rand.Intn(6)+1) // rand.Intn(6) returns from 0 to 5
	
	fmt.Println("--- random float")
	fmt.Println(rand.Float64()) // returne a number between 0 and 1 
	
	fmt.Println("--- random char")
	randNum := rand.Intn(26) // between 0 and 25
	randNum += 97 // randNum is in the ascii values for downcased letters
	fmt.Println( string(randNum) )
	
	fmt.Println("--- random string")
	stringSize := rand.Intn(20)+1 // max size of string is 20
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
	stringBytes := make([]byte, stringSize)
  	for i := range stringBytes {
    	stringBytes[i] = charSet[rand.Intn(len(charSet))]
  	}
	fmt.Println( string(stringBytes) )
}
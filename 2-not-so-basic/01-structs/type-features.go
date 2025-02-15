package main

import(
	"fmt"
	"regexp"
)

// we can create an alias for regular types using type
// it's useful when we want to add functions for a primitive type but scoped
type MyInt int64

func (mi MyInt) ascii()string{
	return string(mi)
}

type Cpf string
func (cpf Cpf) isValid()bool{
	cpfPattern, _ := regexp.Compile("([0-9]{2}[\\.]?[0-9]{3}[\\.]?[0-9]{3}[\\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\\.]?[0-9]{3}[\\.]?[0-9]{3}[-]?[0-9]{2})")
	return cpfPattern.MatchString( string(cpf) )
}

func main(){
	// type as alias
	var exclusiveInteger MyInt = 50
	fmt.Println(exclusiveInteger)
	fmt.Println(exclusiveInteger == 50)
	var regularInt int64 = 50
	// fmt.Println(exclusiveInteger == regularInt) // it doesnt work because they're different types (even in the practice we're the same)
	fmt.Println(exclusiveInteger == MyInt(regularInt) ) // if casting some of them it works
	fmt.Println( int64(exclusiveInteger) == regularInt )

	fmt.Println( exclusiveInteger.ascii() ) // 50 is "2" in ascii
	// fmt.Println( regularInt.ascii() ) // it doesnt work!!

	
	var myCpf Cpf = "123.123.123-12"
	fmt.Println( myCpf.isValid() )
	//fmt.Println( "123.123.123-12".isValid() ) // it doesnt work

}
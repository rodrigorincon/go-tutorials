package main
 
import(
    "fmt"
    "os"
    "reflect"
    "strconv"
)
 
func main() {
    // The first argument is always program name (full path)
    myProgramName := os.Args[0]
    fmt.Println("this progam calls", myProgramName)

    numParams := len(os.Args) - 1
    allParams := os.Args[1:] // the args are separated by space
    fmt.Println(numParams, allParams)

    fmt.Println("Args is type ", reflect.TypeOf(os.Args))
    fmt.Println("The argV is always type ", reflect.TypeOf(os.Args[1]))

    sum := 0
    for _, value := range(allParams){
        convertedVal, _ := strconv.Atoi(value)
        sum += convertedVal
    }
    fmt.Println("the sum is: ", sum)
}
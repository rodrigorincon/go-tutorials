package main
 
import(
    "fmt"
    "strings"
)
 
func main() {
    var name, surname string
    var age int
    var height float32
    
    fmt.Println("What is your name and surname? ")
    fmt.Scan(&name, &surname) // use space to separate what is a var and what is another. 
    // If you pass more than 2 words will consider the next words as answers for next scans 
    // and what's left will be executed in the terminal after this program finishes
    
    fmt.Println("How old are you? ")
    fmt.Scan(&age)
    
    fmt.Println("What is your salary?")
    fmt.Scan(&height)
    
    fmt.Println("The informed data are:\nName: ", name, "\nSurname:", surname, "\nAge: ", age, "\nHeight: ", height)

    // processing input data
    fmt.Println("--------------")

    if(age >= 18){
        fmt.Println("You are legal age")
    }
    timeToRetire := 65 - age
    fmt.Println("You will retire in", timeToRetire)

    fmt.Println( strings.ToUpper(name) )
    fmt.Println( strings.ToLower(surname) )

    if(height < 1.5){
        fmt.Println("you can't go in a rollercoaster")
    }else{
        fmt.Println("you can go in a rollercoaster")
    }
}
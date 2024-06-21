# go-tutorials
Some beggining codes in go to learn the language

### How to install
On Ubuntu, go to the terminal and run `sudo apt-get install golang-go`

To test is everything works, run the command `go version` and see if it returns the downloaded version.

### Running the codes
Enter in the folder where is the file you want to run and execute `go run myfile.go`. This command is good for testing, but take a time to run because it needs to compile before run. When you finish the code, run `go build myfile.go` to compile and create a binary file. To run the binary run `./myfile`

To run a file test, go to its folder and run  `go test FILE_TEST.GO`

### Good Pratices in Go
Variables, constants and functions are camelCase. Public data have the first letter Capitalized and privated first letter lowerCased. An example:
```
// private data (visible only in the package)
cildrenQuantity := 3
fullName := "Jhon Doe" 
func convertToJson(originalFile)

// public data (visible outside the package)
CildrenQuantity := 3
FullName := "Jhon Doe" 
func ConvertToJson(originalFile)
```

### Important notes about the language
Go is: 
- strongly typed
- compiled
- high level language
- multi paradigm (imperative, functional, concurrent, OO (not so much) )
- has garbage collector
- foccused in concurrency
- doesn't have exceptions, inheritance and overload (sobrecarga - métodos com mesmo nome mas params diferentes)
- memory safety (protect against programmer's errors, avoiding memory issues about memory)

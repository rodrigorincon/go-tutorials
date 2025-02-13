# go-tutorials
Some beggining codes in Go to learn the language. All codes here are the first contact with Go and testing possibilities with the language.

The foldes are divides by basic concepts (print in the screen, how to define a var and which var type exists, casing, conditional, loop, functions...), not so basic concepts (structs, pointers, file handling and packages) and intermediate (internet communication, implementing data structures, automated tests and pseudo classes). In `go-differential-things` is explored deeper the possibilities with Go going in its strong points (closures, concepts as defer, implemented a simples https server, real time communication and threads).

### How to install
On Ubuntu, go to the terminal and run `sudo apt-get install golang-go`

To test is everything works, run the command `go version` and see if it returns the downloaded version.

### Running the codes
Enter in the folder where is the file you want to run and execute `go run myfile.go`. This command is good for testing, but take a time to run because it needs to compile before run. When you finish the code, run `go build myfile.go` to compile and create a binary file. To run the binary run `./myfile`

To run a file test, go to its folder and run  `go test FILE_TEST.GO`

### Good Pratices in Go
Variables and functions are camelCase. Public data have the first letter Capitalized and privated first letter lowerCased. But package names are the name of their files and path, so packages is snake_case. Methods usually pass the params as the initials of the class. An example:
```
// private data (visible only in the package)
cildrenQuantity := 3
fullName := "Jhon Doe" 
const USERURL := "localhost"
func convertToJson(originalFile)

// public data (visible outside the package)
CildrenQuantity := 3
const USERURL := "localhost"
FullName := "Jhon Doe" 
func ConvertToJson(originalFile)

// package names
import("path/my_folder/my_file")

// methods
type EmployeeCompany struct{}
func (ec *EmployeeCompany) getName(){}
```

When creates a var, define explicity the type of the var. Example:
```
// Good
var myVar string = someMethod()

//Bad
myVar := someMethod()
```

To comment in Go, always use single line comment (//). The block comment exists ( /* ) but it's used just to comment a broken code that will fixed in the future
```
// Good
// lorem ipsum
// dolor sit amet

//Bad
/* lorem ipsum
   dolor sit amet
*/
```

In Go the dependencies are commited and versioned together the project (vendor folder). It allows you always have everything you need to run your project and doesn't depend of anything. It specially important if you don't have some versioning in their dependencies, because if som of them is updated, your project can break.

### Important notes about the language
Go is: 
- strongly typed
- compiled
- high level language
- multi paradigm (imperative, functional, concurrent)
- doesn't have classes
- has garbage collector
- foccused in concurrency
- doesn't have exceptions, inheritance and overload (sobrecarga - métodos com mesmo nome mas params diferentes)
- memory safety (protect against programmer's errors, avoiding memory issues about memory)

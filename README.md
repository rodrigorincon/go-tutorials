# go-tutorials
Some beggining codes in Go to learn the language. All codes here are the first contact with Go and testing possibilities with the language.

The foldes are divides by basic concepts (print in the screen, how to define a var and which var type exists, casting, conditional, loop, functions...), not so basic concepts (structs, pointers, file handling and packages) and intermediate (internet communication, automated tests, closures, threads and specifity things of Go).

### How to install
On Ubuntu, go to the terminal and run `sudo apt-get install golang-go`

To test is everything works, run the command `go version` and see if it returns the downloaded version.

### Running the codes
Enter in the folder where is the file you want to run and execute `go run myfile.go`. This command is good for testing, but take a time to run because it needs to compile before run. When you finish the code, run `go build myfile.go` to compile and create a binary file. To run the binary run `./myfile`

To run a file test, go to its folder and run  `go test FILE_TEST.GO`

### Good Pratices in Go
Variables and functions are camelCase. Public data have the first letter Capitalized and privated first letter lowerCased. Class methods usually pass the params as the initials of the class. An example:
```
// private data (visible only in the package)
cildrenQuantity := 3
fullName := "Jhon Doe" 
const uSERURL := "localhost"
func convertToJson(originalFile)

// public data (visible outside the package)
CildrenQuantity := 3
const USERURL := "localhost"
FullName := "Jhon Doe" 
func ConvertToJson(originalFile)

// methods
type EmployeeCompany struct{}
func (ec EmployeeCompany) GetName(){}
```

Package names are always the name of folder where it is, so packages is snake_case (with underline). And the name of your packages (and folder's name) must be singular.
```
// GOOD
package my_file
import("path/my_folder/my_file")
my_file.MyFunction()

// BAD
package my-file
import("path/my_folder/my-file")
my-file.MyFunction()

// BAD
package MyFile
import("path/my_folder/MyFile")
MyFile.MyFunction()

// BAD
package users
import("path/my_folder/users")
users.MyFunction()
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

Your module in `go.mod` file must have the url for this project. Use the repository url as the module name for your project.
```
// Good
(in go.mod)
module github.com/username/my_project

(in the code)
import(
   "github.com/username/my_project/my_package"
)

// Bad
(in go.mod)
module my_project

(in the code)
import(
   "my_project/my_package"
)
```

Unit test files are in the same folder that what they test. Other types of test (integration or end-to-end) will be in another folder. When the test file is a unit test (test internal components) use the same package for the code and the test file. When you're doing a integration or end-to-end test, use another package for don't have access to internal structures.
```
// GOOD
(code file)
package my_package

(unit test file)
package my_package

(integration test file)
package my_package_test
```

### Directive structure
A good directory structure to follow is like the example below. You create a go.mod and add your main file (with your main function) inside in `cmd` directory. If you have many entrypoints (many ways to start the project) you can have more than 1 main file, each one with a main function. For it, create a directory inside cmd for each main file. It's good name your file with main function as `main.go`.

The `internal` folder stores your components exclusive for this project. Domain rules, services, database communication, all these things are here.

The `pkg` folder stores your components that can be reusable in another projects. Process queue, database communication (in a lower level) and utils are stored here.

The `build` folder stores files related to build and CI-CD. Docker files and environment configuration files are here.

The `config` folder stores files that read the environment variables and set them in the running project.

The `docs` folder stores all documentation files related to your project. Swagger and diagrams are here.

The `scripts` folder stores the files in another languages and bash files or makefiles.

The `test` folder stores the not-unit tests of your project. High level tests like integration or end-to-end are stored here. Unit test files are always together the files that they test. in `internal` folder.

The `vendor` folder stores the third-party dependencies you use. The libraries listed in you go.mod file are downloaded here. It's used when you're following the strategy to keep all the necessary files to run the project inside the project. In this case, you commit and version the dependencies here. It allows you always have everything you need to run your project and doesn't depend of anything. It specially important if you don't have some versioning in their dependencies, because if some of them is updated, your project can break.

The `website` folder stores assets (images) and web files (javascript, html and cssss) used in your project. If your project have a web interface and renders a html, these files are here.

**OBS: This is not a rule, each project uses more or less folders, joining or splitting some of them.**
```
build/
   Dockerfile
   docker-compose.yml
   start.sh
configs/
cmd/
   api-client/
      main.go
   api-server/
      main.go
docs/
   swagger.json
internal/
   domain/
      person.go
      person_test.go
      car.go
      car_test.go
   service/
      factory_service.go
      factory_service_test.go
      document_validator_service.go
   api/
      bmw_request.go
      bmw_request_test.go
      volkswagen_request.go
pkg/
   aws_util/
     s3.go
   redis_service/
      redis.go 
   utils/
      conversor.go
scripts/
   makefile.sh
test/
   integration/
   e2e/
vendor/
website/
   assets/
   javascript/
   view/
      file.html
go.mod
README.md

```

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

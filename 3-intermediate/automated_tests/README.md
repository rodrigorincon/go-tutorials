# How to test in Go

### Environment for run tests

Your code file and the test file must be in the same directory and use the same package for you have access to private structures. It's specific for unit tests as explained more ahead.

The test file name must ends with `_test.go`.

It's required create a module for test can be run. Execute `go mod init SOME_MODULE_NAME` before run the tests.

### The test file

The test file must use the same package as the file to be tested (specific for unit test).

Import the "testing" library to do the tests. The test functions **must** receive only a single param type `*testing.T`, imported from the library.

If a function named correclty doesn't receive a param `*testing.T` the test will fail.

The test functions must start with `Test` for Go knows what functions it needs to execute.

Functions that don't start with `Test` or after "Test" doesn't come a capital leter will not be executed. E.g. 
```
// GOOD
func TestSomething(t *testing.T) // is executed

// BAD
func Testsomething(t *testing.T) // is not executed

// BAD
func testSomething(t *testing.T) // is not executed

// BAD
func TestSomething() // even run the tests! The command fails
```

The test functions are the most directly test, using if and `t.Error` to log the failed tests. To have a more readable tests and logs, like `expect` or `BeEqualsTo` we need to use some external library.

### Running the tests

Run `go test` to execute all test files in the current folder. To run all tests in inner folders too run `go test ./...`. Run `go test -v` to see more details about the functions.

**OBS:** run `go test TEST_FILE` fails, because it doesn't find the things in the code file.
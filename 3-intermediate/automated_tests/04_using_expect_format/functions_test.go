package some_package

import(
	"testing"
  	"github.com/stretchr/testify/assert" // run "go get github.com/stretchr/testify" to install the library
)

// lib documentation: https://pkg.go.dev/github.com/stretchr/testify@v1.10.0/assert
// lib documentation (mock): https://pkg.go.dev/github.com/stretchr/testify@v1.10.0/mock

func TestStudentPass(t *testing.T){
	grades := []int{3,7,6} // average: 5,33333
	assert.True(t, StudentPass(grades, 5) )
	
	assert.False(t, StudentPass(grades, 6) )
	
	grades = []int{2, 6}
	assert.True(t, StudentPass(grades, 4) )
	
	grades = []int{5}
	assert.True(t, StudentPass(grades, 5) )
	
	grades = []int{}
	assert.False(t, StudentPass(grades, 0) )
}

func TestIsLegalAge(t *testing.T){
	assert.False(t, IsLegalAge(0) )

	assert.False(t, IsLegalAge(17) )
	
	assert.True(t, IsLegalAge(18) )

	assert.True(t, IsLegalAge(19) )
}

func TestSumArr(t *testing.T){
	values := []int{3,7,6} // sum: 16
	assert.Equal(t, SumArr(values), 16)
	
	values = []int{5,9,8} // sum: 22
	assert.Equal(t, SumArr(values), 22)
	
	values = []int{25,31} // sum: 56
	sum := SumArr(values)
	assert.Equal(t, sum, 56, "expected value: %d. Value got: %d", 56, sum) // we can optionally pass a error message to be printed in case of error
}

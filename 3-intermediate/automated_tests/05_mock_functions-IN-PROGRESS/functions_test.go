package some_package

import(
	"testing"
  	"github.com/stretchr/testify/assert" // run "go get github.com/stretchr/testify" to install the library
  	"github.com/stretchr/testify/mock"
)

// lib documentation: https://pkg.go.dev/github.com/stretchr/testify@v1.10.0/assert
// lib documentation (mock): https://pkg.go.dev/github.com/stretchr/testify@v1.10.0/mock
type mockApiFunc struct { 
	mock.Mock 
}
func (m *mockApiFunc) ApiValidateIdNumber(identity string)bool {
	args := m.Called(identity)
    return args.Bool(0)
}

func TestStruct(t *testing.T){
	var m mockApiFunc
	m.On("ApiValidateIdNumber").Return(true)

	person := CreatePerson("123.123.123-12", "jhon")
	assert.Equal(t, person.Id, 123)
	assert.Equal(t, person.Name, "jhon")
	assert.Equal(t, person.IdentityNumber, "123.123.123-12")


	m.On("ApiValidateIdNumber").Return(false)

	person = CreatePerson("123.123.123-12", "jhon")
	assert.Nil(t, person)
}
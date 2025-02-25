package some_package

import(
	"testing"
  	"github.com/stretchr/testify/assert" // run "go get github.com/stretchr/testify" to install the library
  	"github.com/stretchr/testify/mock"
)

// lib documentation: https://pkg.go.dev/github.com/stretchr/testify@v1.10.0/assert
// lib documentation (mock): https://pkg.go.dev/github.com/stretchr/testify@v1.10.0/mock
type mockClass struct { 
	mock.Mock
	Competition
}
// mocking the methods
func (m *mockClass) HasPlayer(player string)bool {
	args := m.Called(player)
    return args.Bool(0) // return the value defined in the test as the response for this method
}
func (m *mockClass) GetPlayerPosition(player string)int {
	args := m.Called(player)
    return args.Int(0) // return the value defined in the test as the response for this method
}

func TestStruct(t *testing.T){
	var m mockClass = mockClass{
		Competition: Competition{
			Name: "my game",
			Prizes: []string{"travel", "book", "cookie"},
		},
	}
	m.On("HasPlayer", "jhon").Return(true) // when runs this method with this param, return this response
	m.On("GetPlayerPosition", "jhon").Return(1) // when runs this method with this param, return this response

	prize, err := m.PrizeAwarded("jhon")
	assert.Nil(t, err)
	assert.Equal(t, prize, "book")

	// second test, testing if jhon is not playing
	m.On("HasPlayer", "jhon").Return(false)
	
	prize, err = m.PrizeAwarded("jhon")
	assert.Error(t, err)
	assert.Equal(err.Error(), "player is not participating of this competition")
	assert.Empty(t, prize) // assert.Empty checks if value is "", nil, false, 0 or []

	// third test, testing if jhon didnt win a prize
	m.On("HasPlayer", "jhon").Return(true)
	m.On("GetPlayerPosition", "jhon").Return(5)
	
	prize, err = m.PrizeAwarded("jhon")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "player doesn't win any prize")
	assert.Empty(t, prize)

	//additional methods
	methodWasCalled := m.AssertCalled(t, "GetPlayerPosition") // check if a method was called. You can pass params to check if the method receives these specifics args
	methodNotCalled := m.AssertNotCalled(t, "GetPlayerPosition") // check if a method was not called. You can pass params to check if the method receives these specifics args
	assert.True(t, methodWasCalled)
	assert.False(t, methodNotCalled)
	methodCalled1Time := m.AssertNumberOfCalls(t, "GetPlayerPosition", 1) // check if a method was called only X times
	assert.False(t, methodCalled1Time)

}
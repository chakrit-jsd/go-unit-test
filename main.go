package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerREpositoryMock{}
	c.On("GetCustomer", 1).Return("John", 20, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	name, age, err := c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)

}

// mock with testify
type CustomerREpositoryMock struct {
	mock.Mock
}

func (m *CustomerREpositoryMock) GetCustomer(id int) (name string, age int, err error) {

	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}

package utils

import "fmt"

type Employee struct {
	Name       string
	Age        int
	Address    string
	IsFullTime bool
}

func (e Employee) GetEmployeeName() {
	fmt.Println(e.Name)
}

// Here we pass copy of the struct type employee in the function
func (e Employee) NewAddress() {
	e.Address = "test@go.dev"
	fmt.Println("Address of this user is", e.Address)
}

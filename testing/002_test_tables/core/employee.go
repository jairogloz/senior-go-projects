package core

import (
	"fmt"
	"strings"
)

// Employee is a struct that represents an employee
type Employee struct {
	Name  string
	Email string
	Age   int
}

// Validate is a method that validates an employee. If the employee is valid
// it returns true, otherwise it returns false and an error.
func Validate(e Employee) (isValid bool, err error) {

	if e.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	if e.Age < 18 {
		return false, fmt.Errorf("age must be greater than 18")
	}

	if strings.Contains(e.Email, "@") == false {
		return false, fmt.Errorf("email must contain @")
	}

	return true, nil
}

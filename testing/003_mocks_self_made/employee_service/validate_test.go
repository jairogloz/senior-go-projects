package employee_service_test

import (
	"errors"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/domain"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/employee_service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type EmployeeStorageMock struct {
	employeeToReturn *domain.Employee
	errToReturn      error
}

func (g *EmployeeStorageMock) GetEmployee(id string) (*domain.Employee, error) {
	return g.employeeToReturn, g.errToReturn
}

func TestService_ValidateEmployee(t *testing.T) {

	employeeStorageMock := &EmployeeStorageMock{}

	service := employee_service.NewService(employeeStorageMock)

	t.Run("valid employee", func(t *testing.T) {

		employeeStorageMock.employeeToReturn = &domain.Employee{
			Status: "active",
		}
		employeeStorageMock.errToReturn = nil

		isValid, err := service.ValidateEmployee("123")
		if isValid != true {
			t.Errorf("expected true, got %v", isValid)
		}
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}

	})

	t.Run("invalid employee", func(t *testing.T) {

		employeeStorageMock.employeeToReturn = &domain.Employee{
			Status: "inactive",
		}
		employeeStorageMock.errToReturn = nil

		isValid, err := service.ValidateEmployee("123")
		assert.False(t, isValid)
		if assert.Error(t, err) {
			assert.Equal(t, "employee with id '123' is not active", err.Error())
		}

	})

	t.Run("storage returns a generic error", func(t *testing.T) {

		employeeStorageMock.employeeToReturn = nil
		employeeStorageMock.errToReturn = errors.New("some error")

		isValid, err := service.ValidateEmployee("123")
		assert.False(t, isValid)
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "error getting employee")
		}

	})

	t.Run("storage returns not found error", func(t *testing.T) {

		employeeStorageMock.employeeToReturn = nil
		employeeStorageMock.errToReturn = domain.ErrNotFound

		isValid, err := service.ValidateEmployee("123")
		assert.False(t, isValid)
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "employee with id '123' not found")
		}
	})

}

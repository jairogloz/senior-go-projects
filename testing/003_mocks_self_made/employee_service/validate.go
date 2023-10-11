package employee_service

import (
	"errors"
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/domain"
)

func (s Service) ValidateEmployee(id string) (bool, error) {
	employee, err := s.storage.GetEmployee(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return false, fmt.Errorf("employee with id '%s' not found", id)
		}
		return false, fmt.Errorf("error getting employee: %s", err.Error())
	}

	if employee.Status != "active" {
		return false, fmt.Errorf("employee with id '%s' is not active", id)
	}

	return true, nil
}

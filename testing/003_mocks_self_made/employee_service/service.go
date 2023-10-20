package employee_service

import "github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/domain"

type Service struct {
	storage domain.EmployeeStorage
}

func NewService(storage domain.EmployeeStorage) domain.EmployeeService {
	if storage == nil {
		panic("storage is nil")
	}
	return &Service{storage: storage}
}

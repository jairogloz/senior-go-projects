package core

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EmployeeStorage interface {
	GetEmployee(id string) (*Employee, error)
	SaveEmployee(employee *Employee) error
}

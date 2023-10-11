package domain

type Employee struct {
	Id        string `json:"_id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Status    string `json:"status" bson:"status"`
}

type EmployeeService interface {
	// ValidateEmployee checks if the status of an employee is "active".
	ValidateEmployee(id string) (bool, error)
}

type EmployeeStorage interface {
	GetEmployee(id string) (*Employee, error)
}

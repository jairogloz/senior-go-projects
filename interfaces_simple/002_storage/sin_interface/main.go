package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Age  int
}

type MySQLStorage struct {
	// ... MySQL database connection details...
}

func (m *MySQLStorage) GetEmployeeByID(id int) (*Employee, error) {
	// ... Code to query the MySQL database for an employee by ID...
	// Mocking retrieved employee for example purpose
	// query := "Select * from employee where..."
	return &Employee{ID: id, Name: "John Doe", Age: 30}, nil
}

func GetEmployee(storage MySQLStorage, id int) (*Employee, error) {
	return storage.GetEmployeeByID(id)
}

func main() {
	mysqlStorage := MySQLStorage{}

	employeeFromMySQL, _ := GetEmployee(mysqlStorage, 1)
	fmt.Printf("Employee from MySQL: %+v\n", employeeFromMySQL)

}

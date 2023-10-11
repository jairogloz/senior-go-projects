package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Age  int
}

type Storage interface {
	GetEmployeeByID(id int) (*Employee, error)
}

type MySQLStorage struct {
	// ... MySQL database connection details...
}

func (m *MySQLStorage) GetEmployeeByID(id int) (*Employee, error) {
	// ... Code to query the MySQL database for an employee by ID...
	// Mocking retrieved employee for example purpose
	return &Employee{ID: id, Name: "John Doe", Age: 30}, nil
}

type MongoDBStorage struct {
	// ... MongoDB database connection details...
}

func (m *MongoDBStorage) GetEmployeeByID(id int) (*Employee, error) {
	// ... Code to query the MongoDB database for an employee by ID...
	// Mocking retrieved employee for example purpose
	return &Employee{ID: id, Name: "Jane Doe", Age: 25}, nil
}

func GetEmployee(storage Storage, id int) (*Employee, error) {
	return storage.GetEmployeeByID(id)
}

func main() {
	mysqlStorage := &MySQLStorage{}
	mongoDBStorage := &MongoDBStorage{}

	employeeFromMySQL, _ := GetEmployee(mysqlStorage, 1)
	fmt.Printf("Employee from MySQL: %+v\n", employeeFromMySQL)

	employeeFromMongoDB, _ := GetEmployee(mongoDBStorage, 1)
	fmt.Printf("Employee from MongoDB: %+v\n", employeeFromMongoDB)
}

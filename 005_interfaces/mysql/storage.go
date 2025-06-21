package mysql

import (
	"database/sql"
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/core"
	_ "github.com/go-sql-driver/mysql"
)

type storage struct {
	db *sql.DB
}

func NewStorage() core.EmployeeStorage {
	// Open a database connection
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database_name")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return &storage{db: db}
}

func (s storage) GetEmployee(id string) (*core.Employee, error) {
	var record core.Employee
	err := s.db.QueryRow("SELECT id, name FROM employees WHERE id=?", id).Scan(&record.ID, &record.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, fmt.Errorf("no rows found")
		}
		return nil, err
	}

	return &record, nil
}

func (s storage) SaveEmployee(employee *core.Employee) error {
	_, err := s.db.Exec("INSERT INTO table_name (name) VALUES (?, ?)", employee.Name)
	if err != nil {
		return err
	}

	return nil
}

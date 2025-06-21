package server

import (
	"encoding/json"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/core"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	storage core.EmployeeStorage
	Router  *httprouter.Router
}

func NewServer(storage core.EmployeeStorage) *Server {

	s := &Server{
		storage: storage,
		Router:  httprouter.New(),
	}

	s.Router.POST("/employees", s.SaveEmployee)
	s.Router.GET("/employees/:id", s.GetEmployee)

	return s
}

func (s *Server) SaveEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse the JSON request body into an Employee struct
	var employee core.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the employee to the database
	err = s.storage.SaveEmployee(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a success response
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetEmployee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	employee, err := s.storage.GetEmployee(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the employee to the response body
	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

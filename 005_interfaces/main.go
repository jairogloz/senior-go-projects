package main

import (
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/mysql"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/server"
	"log"
	"net/http"
)

func main() {

	//mongoStorage := mongodb.NewStorage()
	mysqlStorage := mysql.NewStorage()

	s := server.NewServer(mysqlStorage)

	log.Fatal(http.ListenAndServe(":8080", s.Router))
}

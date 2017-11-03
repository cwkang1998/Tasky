package main

import (
	"logging_go_backend/db"
	// "fmt"
	// "net/http"
)

func main() {
	connection := db.CreateConn()
	connection.AddOwner("Chen Wen Kang", "admin123", "khfy6cwk@nottingham.edu.my")
	connection.GetAllOwners()
	// list := connection.GetLogEntriesList()
	defer connection.CloseConn()
}

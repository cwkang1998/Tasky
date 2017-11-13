package main

import (
	"net/http"
	"projects/tasky_backend/api"
	"projects/tasky_backend/db"
)

//ApiHandler is a custom handler created to enable passing in of DB api

func main() {
	apiHandler := api.ApiHandler{Conn: db.CreateConn()}
	http.HandleFunc("/addTsk", apiHandler.AddTaskEndpoint)
	http.HandleFunc("/getTsks", apiHandler.GetTasksEndpoint)
	http.HandleFunc("/setTskStatus", apiHandler.SetTaskStatusEndpoint)
	http.HandleFunc("/delTsk", apiHandler.DeleteTaskEndpoint)
	http.ListenAndServe("172.17.7.49:8080", nil)
	defer apiHandler.Conn.CloseConn()
}

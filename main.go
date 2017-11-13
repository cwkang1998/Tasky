package main

import (
	"net/http"
	"projects/logging_backend/api"
	"projects/logging_backend/db"
)

//ApiHandler is a custom handler created to enable passing in of DB api

func main() {
	apiHandler := api.ApiHandler{Conn: db.CreateConn()}
	http.HandleFunc("/addProject", apiHandler.AddProjectEndpoint)
	http.HandleFunc("/getProjects", apiHandler.GetProjectsEndpoint)
	http.HandleFunc("/setProjStatus", apiHandler.SetProjStatusEndpoint)
	http.HandleFunc("/addLog", apiHandler.AddLogEndpoint)
	http.HandleFunc("/getLogs", apiHandler.GetLogsEndpoint)
	http.HandleFunc("/addTask", apiHandler.AddTaskEndpoint)
	http.HandleFunc("/getTasks", apiHandler.GetTasksEndpoint)
	http.HandleFunc("/setTaskStatus", apiHandler.SetTaskStatusEndpoint)
	http.ListenAndServe("172.17.7.49:8080", nil)
	defer apiHandler.Conn.CloseConn()
}

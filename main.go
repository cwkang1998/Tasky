package main

import (
	"logging_go_backend/api"
	"logging_go_backend/db"
	"net/http"
)

//ApiHandler is a custom handler created to enable passing in of DB api

func main() {
	apiHandler := api.ApiHandler{Conn: db.CreateConn()}
	http.HandleFunc("/addproject", apiHandler.AddProjectEndpoint)
	http.HandleFunc("/getprojects", apiHandler.GetProjectsEndpoint)
	http.HandleFunc("/addlog", apiHandler.AddLogEndpoint)
	http.HandleFunc("/getlogs", apiHandler.GetLogsEndpoint)
	http.ListenAndServe(":8080", nil)
	defer apiHandler.Conn.CloseConn()
}

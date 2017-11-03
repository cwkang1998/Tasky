package main

import (
	"encoding/json"
	"logging_go_backend/db"
	"logging_go_backend/models"
	"net/http"
)

//ApiHandler is a custom handler created to enable passing in of DB api
type ApiHandler struct {
	Conn *db.Connection
}

//AddProjectEndpoint is an endpoint to add a new project
func (a *ApiHandler) AddProjectEndpoint(w http.ResponseWriter, r *http.Request) {
	//Method Check
	if r.Method == "POST" {
		//Body Check
		if r.Body == nil {
			http.Error(w, "Invalid Request Body", 400)
			return
		}

		//Decoding json request
		var proj models.Project
		err := json.NewDecoder(r.Body).Decode(&proj)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//Json check
		if proj.Title != nil && proj.Description != nil && proj.TimeCreated != nil && proj.OwnerID != nil && proj.Status != nil {
			a.Conn.AddNewProject(&proj)
		} else {
			http.Error(w, "Field Mismatch", 400)
		}

	} else {
		http.Error(w, "Invalid Request", 400)
		return
	}
}

func main() {
	api := ApiHandler{db.CreateConn()}
	http.HandleFunc("/addproject", api.AddProjectEndpoint)
	http.ListenAndServe(":8080", nil)
	defer api.Conn.CloseConn()
}

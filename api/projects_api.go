package api

import (
	"encoding/json"
	"logging_go_backend/models"
	"net/http"
)

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
			return
		}
	} else {
		http.Error(w, "Invalid Request", 400)
	}
}

//GetAllProjectsEndpoint gets all the projects
func (a *ApiHandler) GetAllProjectsEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//Encoding json request
		projectList := a.Conn.GetProjects()
		if len(projectList) > 0 {
			err := json.NewEncoder(w).Encode(projectList)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		}
	} else {
		http.Error(w, "Invalid Request", 400)
	}
}

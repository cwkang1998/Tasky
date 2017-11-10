package api

import (
	"encoding/json"
	"logging_go_backend/models"
	"net/http"
	"strconv"
)

//AddProjectEndpoint is an endpoint to add a new project
func (a *ApiHandler) AddProjectEndpoint(w http.ResponseWriter, r *http.Request) {
	//Cors
	a.corsHandler(&w)
	if r.Method == "OPTIONS" {
		return
	}

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
			err := a.Conn.AddNewProject(proj)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			w.WriteHeader(http.StatusOK)

		} else {
			http.Error(w, "Field Mismatch", 400)
			return
		}
	} else {
		http.Error(w, "Invalid Request", 400)
	}
}

//SetProjStatusEndpoint set the status of a project
func (a *ApiHandler) SetProjStatusEndpoint(w http.ResponseWriter, r *http.Request) {
	//Cors
	a.corsHandler(&w)
	if r.Method == "OPTIONS" {
		return
	}

	//Method Check
	if r.Method == "GET" {
		//Parse form for query values
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error in Form", 400)
			return
		}
		//Look for form value with key "task_id" and status
		id := r.FormValue("project_id")
		status := r.FormValue("status")
		if len(id) != 0 && len(status) != 0 {
			//String to int conversion
			idValue, err1 := strconv.Atoi(id)
			if err1 != nil {
				http.Error(w, "Bad Form Value", 400)
				return
			}
			statusValue, err2 := strconv.Atoi(status)
			if err2 != nil {
				http.Error(w, "Bad Form Value", 400)
				return
			}
			err3 := a.Conn.SetTaskStatus(idValue, statusValue)
			if err3 != nil {
				http.Error(w, "Fail to set project status", 400)
				return
			}
			w.WriteHeader(http.StatusOK)

		} else {
			http.Error(w, "Empty Form Value", 400)
			return
		}
	} else {
		http.Error(w, "Invalid Request", 400)
	}
}

//GetProjectsEndpoint gets all the projects
func (a *ApiHandler) GetProjectsEndpoint(w http.ResponseWriter, r *http.Request) {
	//Cors
	a.corsHandler(&w)
	if r.Method == "OPTIONS" {
		return
	}

	//Method Check
	if r.Method == "GET" {
		//Parse form for query values
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error in Form", 400)
			return
		}
		//Look for form value with key "project_id"
		//If key exist query specific projects, else query all projects
		id := r.FormValue("project_id")
		if len(id) == 0 {
			projectList, err := a.Conn.GetProjects()
			if err != nil {
				http.Error(w, "Fail to get projects", 400)
				return
			}
			if len(projectList) > 0 {
				err := json.NewEncoder(w).Encode(projectList)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
			}
		} else {
			//String to int conversion
			idValue, err1 := strconv.Atoi(id)
			if err1 != nil {
				http.Error(w, "Bad Form Value", 400)
				return
			}
			project, err2 := a.Conn.GetProject(idValue)
			if err2 != nil {
				http.Error(w, "Fail to get project", 400)
				return
			}
			err3 := json.NewEncoder(w).Encode(project)
			if err3 != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}
	} else {
		http.Error(w, "Invalid Request", 400)
	}
}

package api

import (
	"encoding/json"
	"logging_go_backend/models"
	"net/http"
	"strconv"
)

//AddNewOwnerEndpoint is an endpoint to register new owner
func (a *ApiHandler) AddNewOwnerEndpoint(w http.ResponseWriter, r *http.Request) {
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
		var logsEn models.LogEntry

		err := json.NewDecoder(r.Body).Decode(&logsEn)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		//Json check
		if logsEn.Title != nil && logsEn.Description != nil && logsEn.TimeCreated != nil && logsEn.OwnerID != nil && logsEn.ProjectID != nil {
			err := a.Conn.AddNewLogEntry(logsEn)
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

//GetOwnerEndpoint gets the owners
func (a *ApiHandler) GetOwnerEndpoint(w http.ResponseWriter, r *http.Request) {
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
		//Look for form value with key "id"
		//If key exist query specific log entry, else query all projects
		id := r.FormValue("id")
		if len(id) == 0 {
			logEntriesList, err := a.Conn.GetProjects()
			if err != nil {
				http.Error(w, "Fail to get projects", 400)
				return
			}
			if len(logEntriesList) > 0 {
				err := json.NewEncoder(w).Encode(logEntriesList)
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
			logEntry, err2 := a.Conn.GetLogEntry(idValue)
			if err2 != nil {
				http.Error(w, "Fail to get project", 400)
				return
			}
			err3 := json.NewEncoder(w).Encode(logEntry)
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

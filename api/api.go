package api

import (
	"encoding/json"
	"net/http"
	"projects/tasky_backend/db"
	"projects/tasky_backend/models"
	"strconv"
)

type ApiHandler struct {
	Conn *db.Connection
}

func (a *ApiHandler) corsHandler(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding")
}

//AddTaskEndpoint is an endpoint to add a new task
func (a *ApiHandler) AddTaskEndpoint(w http.ResponseWriter, r *http.Request) {
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
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		//Json check
		if task.Description != nil && task.Time != nil {
			err := a.Conn.AddNewTask(task)
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

//SetTaskStatusEndpoint set the status of a task
func (a *ApiHandler) SetTaskStatusEndpoint(w http.ResponseWriter, r *http.Request) {
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
		id := r.FormValue("task_id")
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
				http.Error(w, "Fail to set task status", 400)
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

//GetTasksEndpoint gets the tasks
func (a *ApiHandler) GetTasksEndpoint(w http.ResponseWriter, r *http.Request) {
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
		//Look for form value with key "status"
		//If key exist query, else ignore
		status := r.FormValue("status")
		if len(status) == 0 {
			http.Error(w, "Unspecific Query: Lacking 'Status' Params", 400)
		} else {
			//String to int conversion
			statusValue, err1 := strconv.Atoi(status)
			if err1 != nil {
				http.Error(w, "Bad Form Value", 400)
				return
			}
			task, err2 := a.Conn.GetTasks(statusValue)
			if err2 != nil {
				http.Error(w, "Fail to get tasks", 400)
				return
			}
			err3 := json.NewEncoder(w).Encode(task)
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

// DeleteTaskEndpoint deletes the tasks
func (a *ApiHandler) DeleteTaskEndpoint(w http.ResponseWriter, r *http.Request) {
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
		//Look for form value with key "task_id"
		//If key exist query, else ignore
		id := r.FormValue("task_id")
		if len(id) == 0 {
			http.Error(w, "Unspecific Query: Lacking 'task_id' Params", 400)
		} else {
			//String to int conversion
			idValue, err1 := strconv.Atoi(id)
			if err1 != nil {
				http.Error(w, "Bad Form Value", 400)
				return
			}
			err2 := a.Conn.DeleteTask(idValue)
			if err2 != nil {
				http.Error(w, "Fail to get tasks", 400)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	} else {
		http.Error(w, "Invalid Request", 400)
	}
}

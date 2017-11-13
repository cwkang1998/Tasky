package api

import (
	"encoding/json"
	"net/http"
	"projects/tasky_backend/db"
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
			w.WriteHeader(400)
			w.Write([]byte("Invalid Request Body"))
			return
		}

		//Decoding json request
		var descStruct struct {
			Description *string `json:"description"`
		}

		err := json.NewDecoder(r.Body).Decode(&descStruct)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		//Json check
		if descStruct.Description != nil {
			err := a.Conn.AddNewTask(descStruct.Description)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(400)
			w.Write([]byte("Field Mismatch"))
			return
		}
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Request"))
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
			w.WriteHeader(400)
			w.Write([]byte("Form Error"))
			return
		}
		//Look for form value with key "task_id" and status
		id := r.FormValue("task_id")
		status := r.FormValue("status")
		if len(id) != 0 && len(status) != 0 {
			//String to int conversion
			idValue, err1 := strconv.Atoi(id)
			if err1 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad Form Value"))
				return
			}
			statusValue, err2 := strconv.Atoi(status)
			if err2 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad Form Value"))
				return
			}
			err3 := a.Conn.SetTaskStatus(idValue, statusValue)
			if err3 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Fail to set task status"))
				return
			}
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(400)
			w.Write([]byte("Empty Form Value"))
			return
		}
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Request"))
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
			w.WriteHeader(400)
			w.Write([]byte("Form Error"))
			return
		}
		//Look for form value with key "status"
		//If key exist query, else ignore
		status := r.FormValue("status")
		if len(status) == 0 {
			w.WriteHeader(400)
			w.Write([]byte("Unspecific Query: Lacking 'Status' Params"))
		} else {
			//String to int conversion
			statusValue, err1 := strconv.Atoi(status)
			if err1 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad Form Value"))
				return
			}
			task, err2 := a.Conn.GetTasks(statusValue)
			if err2 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Fail to get tasks"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			err3 := json.NewEncoder(w).Encode(task)
			if err3 != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
		}
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Request"))
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
			w.WriteHeader(400)
			w.Write([]byte("Form Error"))
			return
		}
		//Look for form value with key "task_id"
		//If key exist query, else ignore
		id := r.FormValue("task_id")
		if len(id) == 0 {
			w.WriteHeader(400)
			w.Write([]byte("Unspecific Query: Lacking 'task_id' Params"))
		} else {
			//String to int conversion
			idValue, err1 := strconv.Atoi(id)
			if err1 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad Form Value"))
				return
			}
			err2 := a.Conn.DeleteTask(idValue)
			if err2 != nil {
				w.WriteHeader(400)
				w.Write([]byte("Fail to delete tasks"))
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Request"))
	}
}

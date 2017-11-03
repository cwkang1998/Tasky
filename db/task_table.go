package db

import (
	"fmt"
	"log"
	"logging_go_backend/db/sqlConst"
	"logging_go_backend/models"
)

//AddNewTask creates a new task
func (d *Connection) AddNewTask(timeCreated string, title string, description string, ownerID int, projectID int, logID int, status int) {
	_, err := d.dbInstance.Exec(sqlConst.AddNewTask, timeCreated, title, description, ownerID, projectID, logID, status)
	if err != nil {
		log.Fatalln("Failed to add new task", err)
	}
}

//SetTaskStatus set the status of the task
func (d *Connection) SetTaskStatus(id int, status int) {
	_, err := d.dbInstance.Exec(sqlConst.SetTaskStatus, status, id)
	if err != nil {
		log.Fatalln("Failed to set task status:", err)
	}
}

//GetTasks returns all existing tasks
func (d *Connection) GetTasks() []models.Task {
	rows, err := d.dbInstance.Query(sqlConst.GetTasks)
	if err != nil {
		log.Fatalln("Query Tasks Fail. : ", err)
	}
	var TasksList = make([]models.Task, 1)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID,
			&task.TimeCreated, &task.Title, &task.Description,
			&task.OwnerID, &task.ProjectID, &task.LogID, &task.Status)
		if err != nil {
			fmt.Println("Return Tasks Fail")
			fmt.Println(err)
		}
		TasksList = append(TasksList, task)
	}
	return TasksList
}

//GetTask returns specific task
func (d *Connection) GetTask(id int) *models.Task {
	row := d.dbInstance.QueryRow(sqlConst.GetTask, id)
	var task models.Task
	err := row.Scan(&task.TaskID,
		&task.TimeCreated, &task.Title, &task.Description,
		&task.OwnerID, &task.ProjectID, &task.LogID, &task.Status)
	if err != nil {
		log.Fatalln("Query Task ? Fail. : ", id, err)
	}
	return &task
}

package db

import (
	"logging_go_backend/db/sqlConst"
	"logging_go_backend/models"
)

//AddNewTask creates a new task
func (d *Connection) AddNewTask(tsk models.Task) error {
	_, err := d.dbInstance.Exec(sqlConst.AddNewTask, tsk.TimeCreated, tsk.Title, tsk.Description, tsk.OwnerID, tsk.ProjectID, tsk.LogID, tsk.Status)
	return err
}

//SetTaskStatus set the status of the task
func (d *Connection) SetTaskStatus(id int, status int) error {
	_, err := d.dbInstance.Exec(sqlConst.SetTaskStatus, status, id)
	return err
}

//GetTasks returns all existing tasks
func (d *Connection) GetTasks() ([]models.Task, error) {
	rows, err := d.dbInstance.Query(sqlConst.GetTasks)
	if err != nil {
		return nil, err
	}
	var TasksList = make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID,
			&task.TimeCreated, &task.Title, &task.Description,
			&task.OwnerID, &task.ProjectID, &task.LogID, &task.Status)
		if err != nil {
			return nil, err
		}
		TasksList = append(TasksList, task)
	}
	return TasksList, err
}

//GetTask returns specific task
func (d *Connection) GetTask(id int) (*models.Task, error) {
	row := d.dbInstance.QueryRow(sqlConst.GetTask, id)
	var task models.Task
	err := row.Scan(&task.TaskID,
		&task.TimeCreated, &task.Title, &task.Description,
		&task.OwnerID, &task.ProjectID, &task.LogID, &task.Status)
	return &task, err
}

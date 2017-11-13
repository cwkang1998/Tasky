package db

import (
	"database/sql"
	"log"
	"projects/logging_backend/db/sqlConst"
	"projects/logging_backend/models"

	//Sqlite3 Driver import
	_ "github.com/mattn/go-sqlite3"
)

//Connection is a database connection instance to talk to database
type Connection struct {
	dbInstance *sql.DB
}

func tryInitDatabase(db *sql.DB) {
	_, err := db.Exec(sqlConst.TaskTableCreate)
	if err != nil {
		panic(err)
	}
}

//CreateConn generates a new Connection instance to the caller
func CreateConn() *Connection {
	dbInstance, err := sql.Open("sqlite3", "./db/sqlite3/go_backend.db")
	if err != nil {
		log.Fatal("Fail to establish connection")
	} else {
		tryInitDatabase(dbInstance)
	}
	dbConn := Connection{dbInstance}
	return &dbConn
}

//CloseConn used to close the database connection. Called at defer.
func (d *Connection) CloseConn() {
	err := d.dbInstance.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//AddNewTask creates a new task
func (d *Connection) AddNewTask(tsk models.Task) error {
	_, err := d.dbInstance.Exec(AddNewTask, tsk.TimeCreated, tsk.Description)
	return err
}

//SetTaskStatus set the status of the task
func (d *Connection) SetTaskStatus(id int, time string, status int) error {
	_, err := d.dbInstance.Exec(SetTaskStatus, status, time, id)
	return err
}

//DeleteTask deletes a task
func (d *Connection) DeleteTask(id int) error {
	_, err := d.dbInstance.Exec(DeleteTask, id)
	return err
}

//GetTasks returns all existing tasks
func (d *Connection) GetTasks(status int) ([]models.Task, error) {
	rows, err := d.dbInstance.Query(GetTasks, status)
	if err != nil {
		return nil, err
	}
	var TasksList = make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID,
			&task.TimeCreated, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}
		TasksList = append(TasksList, task)
	}
	return TasksList, err
}

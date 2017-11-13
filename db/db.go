package db

import (
	"database/sql"
	"log"
	"projects/tasky_backend/models"
	"time"

	//Sqlite3 Driver import
	_ "github.com/mattn/go-sqlite3"
)

//Connection is a database connection instance to talk to database
type Connection struct {
	dbInstance *sql.DB
}

func tryInitDatabase(db *sql.DB) {
	_, err := db.Exec(TaskTableCreate)
	if err != nil {
		panic(err)
	}
}

//CreateConn generates a new Connection instance to the caller
func CreateConn() *Connection {
	dbInstance, err := sql.Open("sqlite3", "./db/tasky_app.db")
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
func (d *Connection) AddNewTask(description *string) error {
	_, err := d.dbInstance.Exec(AddNewTask, time.Now().Format("2006-01-02"), *description)
	return err
}

//Future possible implementation
// func (d *Connection) AddNewTask(tsk models.Task) (*models.Task, error) {
// 	_, err := d.dbInstance.Exec(AddNewTask, time.Now().Format("yyyy-MM-dd"), tsk.Description)
// 	if err != nil {
// 		return &models.Task{}, err
// 	}
// 	row := d.dbInstance.QueryRow(QueryNewTask)
// 	var task models.Task
// 	err1 := row.Scan(&task.TaskID, &task.Time, &task.Description, &task.Status)
// 	if err1 != nil {
// 		return &models.Task{}, err1
// 	}
// 	return &task, err1
// }

//SetTaskStatus set the status of the task
func (d *Connection) SetTaskStatus(id int, status int) error {
	_, err := d.dbInstance.Exec(SetTaskStatus, status, time.Now().Format("yyyy-MM-dd"), id)
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
			&task.Time, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}
		TasksList = append(TasksList, task)
	}
	return TasksList, err
}

package db

import (
	"database/sql"
	"log"
	"projects/logging_backend/db/sqlConst"

	//Sqlite3 Driver import
	_ "github.com/mattn/go-sqlite3"
)

//Connection is a database connection instance to talk to database
type Connection struct {
	dbInstance *sql.DB
}

func tryInitDatabase(db *sql.DB) {
	_, err := db.Exec(sqlConst.LogTableCreate)
	if err != nil {
		panic(err)
	}
	_, err1 := db.Exec(sqlConst.TaskTableCreate)
	if err1 != nil {
		panic(err1)
	}
	_, err2 := db.Exec(sqlConst.ProjectsTableCreate)
	if err2 != nil {
		panic(err2)
	}
	_, err3 := db.Exec(sqlConst.OwnersTableCreate)
	if err3 != nil {
		panic(err3)
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

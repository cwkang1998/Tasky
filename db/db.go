package db

import (
	"database/sql"
	"fmt"
	"log"
	"logging_go_backend/db/sqlConst"
	"logging_go_backend/models"

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

//GetAllOwners returns all existing owners
func (d *Connection) GetAllOwners() {
	queryString := "SELECT owner_id,owner_name,owner_email,last_login_time FROM owners;"
	rows, err := d.dbInstance.Query(queryString)
	if err != nil {
		log.Fatalln("Query failed.")
		log.Fatalln(err)
	}
	for rows.Next() {
		var owner models.Owners
		err := rows.Scan(&owner.OwnerID, &owner.OwnerName, &owner.OwnerEmail, &owner.LastLogin)
		if err != nil {
			fmt.Println("Scanning failed.")
			fmt.Println(err)
		}
		fmt.Println(owner)
	}
}

//AddOwner registers the user into database
func (d *Connection) AddOwner(name, password, email string) {
	queryString := "INSERT INTO owners(owner_name, owner_password, owner_email) VALUES (?,?,?);"
	_, err := d.dbInstance.Exec(queryString, name, password, email)
	if err != nil {
		fmt.Println("Adding new owner failed: ", err)
	}
}

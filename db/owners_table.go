package db

import (
	"fmt"
	"log"
	"logging_go_backend/models"
)

//GetAllOwners returns all existing owners
func (d *Connection) GetAllOwners() {
	queryString := "SELECT owner_id,owner_name,owner_email,last_login_time FROM owners;"
	rows, err := d.dbInstance.Query(queryString)
	if err != nil {
		log.Fatalln("Query Owners failed.")
		log.Fatalln(err)
	}
	for rows.Next() {
		var owner models.Owners
		err := rows.Scan(&owner.OwnerID, &owner.OwnerName, &owner.OwnerEmail, &owner.LastLogin)
		if err != nil {
			fmt.Println("Returning Owners failed.")
			fmt.Println(err)
		}
		fmt.Println(*owner.OwnerName)
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

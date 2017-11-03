package db

import (
	"fmt"
	"log"
	"logging_go_backend/db/sqlConst"
	"logging_go_backend/models"
)

//AddNewProject creates a new project
func (d *Connection) AddNewProject(timeCreated string, title string, description string, ownerID int, status int) {
	_, err := d.dbInstance.Exec(sqlConst.AddNewProject, timeCreated, title, description, ownerID, status)
	if err != nil {
		log.Fatalln("Failed to add new project", err)
	}
}

//SetProjectStatus set the status of the project
func (d *Connection) SetProjectStatus(id int, status int) {
	_, err := d.dbInstance.Exec(sqlConst.SetProjectStatus, status, id)
	if err != nil {
		log.Fatalln("Failed to set project status:", err)
	}
}

//GetProjects returns all existing project
func (d *Connection) GetProjects() []models.Project {
	rows, err := d.dbInstance.Query(sqlConst.GetProjects)
	if err != nil {
		log.Fatalln("Query Project Fail. : ", err)
	}
	var ProjList = make([]models.Project, 1)
	for rows.Next() {
		var proj models.Project
		err := rows.Scan(&proj.ProjectID,
			&proj.TimeCreated, &proj.Title, &proj.Description,
			&proj.OwnerID, &proj.Status)
		if err != nil {
			fmt.Println("Return Project Fail")
			fmt.Println(err)
		}
		ProjList = append(ProjList, proj)
	}
	return ProjList
}

//GetProject returns specific project
func (d *Connection) GetProject(id int) *models.Project {
	row := d.dbInstance.QueryRow(sqlConst.GetProject, id)
	var proj models.Project
	err := row.Scan(&proj.ProjectID,
		&proj.TimeCreated, &proj.Title, &proj.Description,
		&proj.OwnerID, &proj.Status)
	if err != nil {
		log.Fatalln("Query project ? Fail. : ", id, err)
	}
	return &proj
}

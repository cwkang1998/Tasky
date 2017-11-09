package db

import (
	"log"
	"logging_go_backend/db/sqlConst"
	"logging_go_backend/models"
)

//AddNewProject creates a new project
func (d *Connection) AddNewProject(proj models.Project) error {
	_, err := d.dbInstance.Exec(sqlConst.AddNewProject, proj.TimeCreated, proj.Title, proj.Description, proj.OwnerID, proj.Status)

	return err
}

//SetProjectStatus set the status of the project
func (d *Connection) SetProjectStatus(id int, status int) error {
	_, err := d.dbInstance.Exec(sqlConst.SetProjectStatus, status, id)

	return err
}

//GetProjects returns all existing project
func (d *Connection) GetProjects() ([]models.Project, error) {
	rows, err := d.dbInstance.Query(sqlConst.GetProjects)
	if err != nil {
		return nil, err
	}
	var ProjList = make([]models.Project, 0)
	for rows.Next() {
		var proj models.Project
		err := rows.Scan(&proj.ProjectID,
			&proj.TimeCreated, &proj.Title, &proj.Description,
			&proj.OwnerID, &proj.Status)
		//Error handling
		if err != nil {
			log.Fatal(err)
		}
		ProjList = append(ProjList, proj)

	}
	return ProjList, err
}

//GetProject returns specific project
func (d *Connection) GetProject(id int) (*models.Project, error) {
	row := d.dbInstance.QueryRow(sqlConst.GetProject, id)
	var proj models.Project
	err := row.Scan(&proj.ProjectID,
		&proj.TimeCreated, &proj.Title, &proj.Description,
		&proj.OwnerID, &proj.Status)
	return &proj, err
}

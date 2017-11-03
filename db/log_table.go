package db

//log_table manipulates the log_entry table

import (
	"fmt"
	"log"
	"logging_go_backend/db/sqlConst"
	"logging_go_backend/models"
)

//AddNewLogEntry creates a new entry in the entry table
//However it DOES NOT include the adding of new task
func (d *Connection) AddNewLogEntry(logentry models.LogEntry) {
	_, err := d.dbInstance.Exec(sqlConst.AddNewLogEntry, logentry.TimeCreated, logentry.Title, logentry.Description, logentry.OwnerID, logentry.ProjectID)
	if err != nil {
		log.Fatalln("Failed to add new log entry:", err)
	}
}

//GetLogEntriesList returns all existing log entries in a summary
func (d *Connection) GetLogEntriesList() []models.LogEntry {
	rows, err := d.dbInstance.Query(sqlConst.GetLogEntries)
	if err != nil {
		log.Fatalln("Query Log Entries Fail. : ", err)
	}
	var Loglist = make([]models.LogEntry, 1)
	for rows.Next() {
		var LogSumm models.LogEntry
		err := rows.Scan(&LogSumm.LogID,
			&LogSumm.TimeCreated, &LogSumm.Title, &LogSumm.Description,
			&LogSumm.OwnerID, &LogSumm.ProjectID)
		if err != nil {
			fmt.Println("Return Log Entries Fail")
			fmt.Println(err)
		}
		Loglist = append(Loglist, LogSumm)
	}
	return Loglist
}

//GetLogEntry returns specific log entry
func (d *Connection) GetLogEntry(id int) *models.LogEntry {
	row := d.dbInstance.QueryRow(sqlConst.GetLogEntry, id)
	var LogDetails models.LogEntry
	err := row.Scan(&LogDetails.LogID,
		&LogDetails.TimeCreated, &LogDetails.Title, &LogDetails.Description,
		&LogDetails.OwnerID, &LogDetails.ProjectID)
	if err != nil {
		log.Fatalln("Query Log Entry ? Fail. : ", id, err)
	}
	return &LogDetails
}

package db

//log_table manipulates the log_entry table

import (
	"projects/logging_backend/db/sqlConst"
	"projects/logging_backend/models"
)

//AddNewLogEntry creates a new entry in the entry table
//However it DOES NOT include the adding of new task
func (d *Connection) AddNewLogEntry(logentry models.LogEntry) error {
	_, err := d.dbInstance.Exec(sqlConst.AddNewLogEntry, logentry.TimeCreated, logentry.Title, logentry.Description, logentry.OwnerID, logentry.ProjectID)
	return err
}

//GetLogEntriesList returns all existing log entries in a summary
func (d *Connection) GetLogEntriesList() ([]models.LogEntry, error) {
	rows, err := d.dbInstance.Query(sqlConst.GetLogEntries)
	if err != nil {
		return nil, err
	}
	var Loglist = make([]models.LogEntry, 0)
	for rows.Next() {
		var LogSumm models.LogEntry
		err := rows.Scan(&LogSumm.LogID,
			&LogSumm.TimeCreated, &LogSumm.Title, &LogSumm.Description,
			&LogSumm.OwnerID, &LogSumm.ProjectID)
		if err != nil {
			return nil, err
		}
		Loglist = append(Loglist, LogSumm)
	}
	return Loglist, err
}

//GetLogEntry returns specific log entry
func (d *Connection) GetLogEntry(id int) (*models.LogEntry, error) {
	row := d.dbInstance.QueryRow(sqlConst.GetLogEntry, id)
	var LogDetails models.LogEntry
	err := row.Scan(&LogDetails.LogID,
		&LogDetails.TimeCreated, &LogDetails.Title, &LogDetails.Description,
		&LogDetails.OwnerID, &LogDetails.ProjectID)
	return &LogDetails, err
}

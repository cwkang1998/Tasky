package sqlConst

const AddNewLogEntry string = `
INSERT INTO 
 log_entries(time_created, title, description, owner_id, project_id)
 VALUES(?,?,?,?,?)`

const GetLogEntries string = `
SELECT log_id, time_created,
 title, owner_id, project_id
 FROM log_entries`

const GetLogEntry string = `
 SELECT log_id, time_created, title,
 description, owner_id, project_id
 FROM log_entries WHERE log_id = ?`

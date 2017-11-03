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

const AddNewTask string = `
INSERT INTO 
tasks(time_created, title, description, 
	owner_id, project_id, log_id, status)
VALUES(?,?,?,?,?,?,?)`

const SetTaskStatus string = `
UPDATE tasks SET status = ? WHERE task_id = ?`

const GetTasks string = `
SELECT task_id, time_created, title,
description, owner_id, project_id, log_id, status
FROM tasks`

const GetTask string = `
SELECT task_id, time_created, title,
description, owner_id, project_id, log_id, status
FROM tasks WHERE task_id = ?`

const AddNewProject string = `
INSERT INTO 
projects(time_created, title, description, 
	owner_id, status)
VALUES(?,?,?,?,?)`

const SetProjectStatus string = `
UPDATE projects SET status = ? WHERE project_id = ?`

const GetProjects string = `
SELECT project_id, time_created, title,
description, owner_id, status
FROM projects`

const GetProject string = `
SELECT project_id, time_created, title,
description, owner_id, status
FROM projects WHERE project_id = ?`

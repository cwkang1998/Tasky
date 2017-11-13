package db

//SQL for creating table

const TaskTableCreate string = `
CREATE TABLE IF NOT EXISTS tasks(
	task_id INTEGER PRIMARY KEY,
	time TEXT NOT NULL,
	description TEXT NOT NULL,
	status INTEGER NOT NULL
	);`

const AddNewTask string = `
INSERT INTO 
tasks(time, description, status)
VALUES(?,?,0);`

const QueryNewTask string = `
SELECT * FROM tasks WHERE status = 0 ORDER BY id DESC LIMIT 1;
`

const SetTaskStatus string = `
UPDATE tasks SET status = ?, time = ? WHERE task_id = ?;`

const GetTasks string = `
SELECT task_id, time,
description,status FROM tasks WHERE status = ?;`

const DeleteTask string = `
DELETE FROM tasks WHERE task_id = ?;`

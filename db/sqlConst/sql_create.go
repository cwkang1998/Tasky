package sqlConst

//SQL for creating table

const LogTableCreate string = `
CREATE TABLE IF NOT EXISTS log_entries(
	log_id INTEGER PRIMARY KEY,
	time_created TEXT NOT NULL,
	title TEXT NOT NULL ,
	description TEXT NOT NULL,
	owner_id INTEGER NOT NULL,
	project_id INTEGER NOT NULL,
	FOREIGN KEY(owner_id) REFERENCES owners(owner_id),
	FOREIGN KEY(project_id) REFERENCES projects(project_id)
	);`

const TaskTableCreate string = `
CREATE TABLE IF NOT EXISTS tasks(
	task_id INTEGER PRIMARY KEY,
	time_created TEXT NOT NULL,
	title TEXT NOT NULL ,
	description TEXT NOT NULL,
	owner_id INTEGER NOT NULL,
	project_id INTEGER NOT NULL,
	log_id INTEGER NOT NULL,
	status INTEGER NOT NULL,
	FOREIGN KEY(owner_id) REFERENCES owner(owner_id),
	FOREIGN KEY(project_id) REFERENCES projects(project_id),
	FOREIGN KEY(log_id) REFERENCES log_entries(log_id)
	);`

const ProjectsTableCreate string = `
CREATE TABLE IF NOT EXISTS projects(
	project_id INTEGER PRIMARY KEY,
	time_created TEXT NOT NULL,
	title TEXT NOT NULL ,
	description TEXT NOT NULL,
	owner_id INTEGER NOT NULL,
	FOREIGN KEY(owner_id) REFERENCES owners(owner_id)
	);`

const OwnersTableCreate string = `
CREATE TABLE IF NOT EXISTS owners(
	owner_id INTEGER PRIMARY KEY,
	owner_name TEXT NOT NULL,
	owner_password TEXT,
	owner_email TEXT NOT NULL UNIQUE,
	token TEXT,
	last_login_time TEXT
	);`

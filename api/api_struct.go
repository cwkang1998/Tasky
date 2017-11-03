package api

import "logging_go_backend/db"

type ApiHandler struct {
	Conn *db.Connection
}

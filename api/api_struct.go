package api

import (
	"logging_go_backend/db"
	"net/http"
)

type ApiHandler struct {
	Conn *db.Connection
}

func (a *ApiHandler) corsHandler(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}

package main

import (
	"net/http"
	"projects/tasky_app/api"
	"projects/tasky_app/db"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/www/index.html")
}

func main() {
	// Provide files for frontend
	http.Handle("/node_modules/", http.StripPrefix("/node_modules", http.FileServer(http.Dir("./ui/node_modules"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./ui/www/js"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./ui/www/css"))))

	// Apis Hosting
	apiHandler := api.ApiHandler{Conn: db.CreateConn()}
	http.HandleFunc("/addTsk", apiHandler.AddTaskEndpoint)
	http.HandleFunc("/getTsks", apiHandler.GetTasksEndpoint)
	http.HandleFunc("/setTskStatus", apiHandler.SetTaskStatusEndpoint)
	http.HandleFunc("/delTsk", apiHandler.DeleteTaskEndpoint)
	http.HandleFunc("/", indexHandler)

	// Host app frontend
	http.ListenAndServe("localhost:8080", nil)

	// Close connection on app end
	defer apiHandler.Conn.CloseConn()
}

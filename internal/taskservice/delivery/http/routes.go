package http

import (
	"net/http"
)

func SetupUserRoutes() {
	http.HandleFunc("/check", index)
	http.HandleFunc("/create", createTask)
}

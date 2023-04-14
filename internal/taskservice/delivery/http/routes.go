package http

import (
	"net/http"
)

func SetupUserRoutes() {
	http.HandleFunc("/", index)
}

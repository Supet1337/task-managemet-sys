package main

import (
	routes "golang-pkg/internal/taskservice/delivery/http"
	"net/http"
	//"html/template"
)

func main() {
	routes.SetupUserRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error")
	}
}

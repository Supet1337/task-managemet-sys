package http

import (
	"encoding/json"
	"fmt"
	models "golang-pkg/internal/taskservice"
	eventLogic "golang-pkg/internal/taskservice/usecase"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fmt.Fprintf(w, "HELLO POST")
	} else if r.Method == "GET" {
		fmt.Fprintf(w, "HELLO GET")
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task models.Task

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		log.Println(err)
	}
	//{
	//	"taskId": 1,
	//	"status": "STAT",
	//	"name": "NAME",
	//	"description": "DESC",
	//	"timeStart": "2023-01-02T15:04:05Z",
	//	"timeEnd": "2023-02-02T15:04:05Z"
	//}

	err = eventLogic.CreateTask(task)

	jsonTask, err := json.Marshal(task)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonTask)

}

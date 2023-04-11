package main

import (
	"fmt"
	"net/http"
	//"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")
}

func handleFunc() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error")
	}
}

func main() {

	handleFunc()
}

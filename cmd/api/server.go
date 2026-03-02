package main

import (
	"fmt"
	"net/http"
)

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello here is the /Get Method"))

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error in form", 404)
			return
		}
		fmt.Println(r.Form)
	case http.MethodPut:
		w.Write([]byte("Hello here is the Put Method"))
	case http.MethodDelete:
		w.Write([]byte("Hello here is the Delete Method"))
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are on root route"))
}

func main() {
	port := "localhost:8080"

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/teachers", teachersHandler)

	fmt.Println("server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error is server starting err: ", err)
	}
}

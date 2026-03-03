package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	mw "restapi/inrernal/api/middlewares"
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

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers", teachersHandler)

	server := &http.Server{
		Addr:      port,
		Handler:   mw.Compression(mw.SecurityHedders((mw.Cors(mux)))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("server is running on https://localhost:8080")
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		fmt.Println("Error in server starting err: ", err)
	}
}

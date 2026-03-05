package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/inrernal/api/middlewares"
	"time"
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

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	secureMux := applyMiddlewares(
		mux,
		mw.Hpp(hppOptions),
		mw.Compression,
		mw.SecurityHeaders,
		// mw.ResponseTimeMiddleware,
		rl.Middleware,
		mw.Cors,
	)
	// Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}

type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

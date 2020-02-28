package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello World!"))
}

func main() {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", RootEndpoint).Methods("GET")

	http.ListenAndServe(":12345", handlers.CORS(headers, methods, origins)(router))
}

package main

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	host := "127.0.0.1"
	port := 4298

	server := core.InitializeCore()

	router := mux.NewRouter()

	router.Handle("/add/{schemaname}", handlers.NewSchemaAddHandler(server)).Methods("POST")
	router.Handle("/update/{schemaname}", handlers.NewSchemaUpdateHandler(server)).Methods("POST")
	router.Handle("/remove/{schemaname}", handlers.NewSchemaRemoveHandler(server)).Methods("POST")
	router.Handle("/validate/{schemaname}", handlers.NewSchemaValidateHandler(server)).Methods("POST")

	log.Println(fmt.Sprintf("Starting at %v:%d", host, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), router))
}

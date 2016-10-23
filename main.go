package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/Enapiuz/SchemaStorage/handlers"
	"github.com/gorilla/mux"
)

func main() {
	host := "127.0.0.1"
	port := 4298

	router := mux.NewRouter()

	router.Handle("/schema/add/{schemaname}", handlers.NewSchemaAddHandler()).Methods("POST")
	router.Handle("/schema/validate/{schemaname}", handlers.NewSchemaValidateHandler()).Methods("POST")
	router.Handle("/schema/remove/{schemaname}", handlers.NewSchemaRemoveHandler()).Methods("POST")

	log.Println(fmt.Sprintf("Starting at %v:%d", host, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), router))
}

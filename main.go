package main

import (
	"flag"
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var host, port, mongoHost, mongoPort string
	flag.StringVar(&host, "sh", "127.0.0.1", "Server host")
	flag.StringVar(&port, "sp", "4298", "Server port")
	flag.StringVar(&mongoHost, "mh", "127.0.0.1", "MongoDB host")
	flag.StringVar(&mongoPort, "mp", "27017", "MongoDB port")
	flag.Parse()

	server := core.InitializeCore(mongoHost, mongoPort)
	router := mux.NewRouter()

	router.Handle("/add/{schemaname}", handlers.NewSchemaAddHandler(server)).Methods("POST")
	router.Handle("/update/{schemaname}", handlers.NewSchemaUpdateHandler(server)).Methods("POST")
	router.Handle("/remove/{schemaname}", handlers.NewSchemaRemoveHandler(server)).Methods("POST")
	router.Handle("/validate/{schemaname}", handlers.NewSchemaValidateHandler(server)).Methods("POST")

	log.Println(fmt.Sprintf("Starting at %v:%d", host, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), router))
}

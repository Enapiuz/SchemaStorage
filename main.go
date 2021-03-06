package main

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"flag"
)

func main() {
	var host, port, mongoHost, mongoPort string
	flag.StringVar(&host, "sh", "127.0.0.1", "Server host")
	flag.StringVar(&port, "sp", "4298", "Server port")
	flag.StringVar(&mongoHost, "mh", "localhost", "MongoDB host")
	flag.StringVar(&mongoPort, "mp", "32771", "MongoDB port")
	flag.Parse()

	server := core.InitializeCore(mongoHost, mongoPort)
	router := mux.NewRouter()

	router.Handle("/add/{schemaname}", handlers.NewSchemaAddHandler(server)).Methods("POST")
	router.Handle("/update/{schemaname}", handlers.NewSchemaUpdateHandler(server)).Methods("POST")
	router.Handle("/remove/{schemaname}", handlers.NewSchemaRemoveHandler(server)).Methods("POST")
	router.Handle("/validate/{schemaname}", handlers.NewSchemaValidateHandler(server)).Methods("POST")

	log.Println(fmt.Sprintf("Starting at %v:%v", host, port))
	log.Fatal(http.ListenAndServe(host + ":" + port, router))
}

package main

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/core"
	"github.com/Enapiuz/SchemaStorage/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	host := "127.0.0.1"
	port := 4298

	server := initializeCore()

	router := mux.NewRouter()

	router.Handle("/schema/add/{schemaname}", handlers.NewSchemaAddHandler(server)).Methods("POST")
	router.Handle("/schema/validate/{schemaname}", handlers.NewSchemaValidateHandler(server)).Methods("POST")
	router.Handle("/schema/remove/{schemaname}", handlers.NewSchemaRemoveHandler(server)).Methods("POST")

	log.Println(fmt.Sprintf("Starting at %v:%d", host, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), router))
}

func initializeCore() *core.Core {
	mongo, err := mgo.Dial("localhost:32771")
	if err != nil {
		panic(err)
	}

	newCore := core.NewCore(mongo)
	index := mgo.Index{
		Key: []string{"name"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	newCore.GetCollection(core.SchemaCollection).EnsureIndex(index)

	return newCore
}

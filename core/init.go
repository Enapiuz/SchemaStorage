package core

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func InitializeCore() *Core {
	mongo, err := mgo.Dial("localhost:32771")
	if err != nil {
		panic(err)
	}

	newCore := NewCore(mongo)
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = newCore.GetCollection(SchemaCollection).EnsureIndex(index)

	if err != nil {
		panic(fmt.Sprintf("Can't create index on collection %s", SchemaCollection))
	}

	return newCore
}

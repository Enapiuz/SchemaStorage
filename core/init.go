package core

import (
	"fmt"
	"github.com/Enapiuz/SchemaStorage/repository"
	"gopkg.in/mgo.v2"
)

func InitializeCore(mongoHost, mongoPort string) *Core {
	mongo, err := mgo.Dial(fmt.Sprintf("%v:%v", mongoHost, mongoPort))
	if err != nil {
		panic(err)
	}

	database := mongo.DB(SchemaStorage)
	repo := repository.NewRepository(database, SchemaCollection)
	newCore := NewCore(repo)

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = newCore.Repo.GetCollection(SchemaCollection).EnsureIndex(index)

	if err != nil {
		panic(fmt.Sprintf("Can't create index on collection %s", SchemaCollection))
	}

	return newCore
}

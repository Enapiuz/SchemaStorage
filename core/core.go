package core

import "gopkg.in/mgo.v2"

const (
	SchemaCollection = "schemes"
)

type Core struct {
	Mongo *mgo.Session
}

func NewCore(mongo *mgo.Session) *Core {
	core := Core{Mongo: mongo}
	return &core
}

func (c *Core) GetDB() *mgo.Database {
	return c.Mongo.DB("schema_storage")
}

func (c *Core) GetCollection(Name string) *mgo.Collection {
	return c.GetDB().C(Name)
}

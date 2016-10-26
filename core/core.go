package core

import "gopkg.in/mgo.v2"

const (
	SchemaCollection = "schemes"
	SchemaStorage    = "schema_storage"
)

type Core struct {
	Mongo *mgo.Session
}

func NewCore(mongo *mgo.Session) *Core {
	core := Core{Mongo: mongo}
	return &core
}

func (c *Core) GetDB() *mgo.Database {
	return c.Mongo.DB(SchemaStorage)
}

func (c *Core) GetCollection(Name string) *mgo.Collection {
	return c.GetDB().C(Name)
}

package core

import "gopkg.in/mgo.v2"

type Core struct {
	Mongo *mgo.Session
}

func NewCore(mongo *mgo.Session) *Core {
	core := Core{Mongo: mongo}
	return &core
}

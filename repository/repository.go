package repository

import (
	"gopkg.in/mgo.v2"
)

type Repository struct {
	db             *mgo.Database
	collectionName string
}

func NewRepository(db *mgo.Database, collectionName string) *Repository {
	return &Repository{db: db, collectionName: collectionName}
}

func (r *Repository) GetCollection(name string) *mgo.Collection {
	return r.db.C(name)
}

func (r *Repository) GetSchemaCollection() *mgo.Collection {
	return r.GetCollection(r.collectionName)
}

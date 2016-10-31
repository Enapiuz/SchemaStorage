package repository

import (
	"github.com/Enapiuz/SchemaStorage/db_models"
)

func (r *Repository) InsertSchema(newSchema *db_models.Schema) error {
	collection := r.GetSchemaCollection()
	err := collection.Insert(newSchema)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteSchema(schemaName string) error {
	collection := r.GetSchemaCollection()
	err := collection.Remove(struct{ Name string }{Name: schemaName})
	return err
}

func (r *Repository) UpdateSchema(schemaName string, newSchema *db_models.Schema) error {
	collection := r.GetSchemaCollection()
	err := collection.Update(struct{ Name string }{Name: schemaName}, newSchema)
	return err
}

func (r *Repository) GetSchemaByName(schemaName string) (*db_models.Schema, error) {
	collection := r.GetSchemaCollection()
	var foundSchema db_models.Schema
	err := collection.Find(struct {
		Name string
	}{Name: schemaName}).One(&foundSchema)
	if err != nil {
		return nil, err
	}
	return &foundSchema, nil
}

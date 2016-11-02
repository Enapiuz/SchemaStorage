package db_models

type Schema struct {
	Name    string
	Data    string
	Version int
}

func NewSchema(name string, data string) *Schema {
	return &Schema{Name: name, Data: data}
}

package db_models

type Schema struct {
	Name string
	Data string
}

func NewSchema(name string, data string) *Schema {
	return &Schema{Name: name, Data: data}
}

package core

import (
	"github.com/Enapiuz/SchemaStorage/repository"
)

const (
	SchemaCollection = "schemes"
	SchemaStorage    = "schema_storage"
)

type Core struct {
	Repo *repository.Repository
}

func NewCore(repo *repository.Repository) *Core {
	core := Core{Repo: repo}
	return &core
}

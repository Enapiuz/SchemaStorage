package validators

import (
	"bytes"
	"github.com/Enapiuz/SchemaStorage/repository"
	"github.com/lestrrat/go-jsschema"
	"github.com/lestrrat/go-jsval/builder"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"log"
)

func ValidateJSONBytes(schemaBytes *[]byte) error {
	s, err := schema.Read(bytes.NewBuffer(*schemaBytes))
	if err != nil {
		log.Printf("Invalid schema: %s", err)
		return err
	}

	b := builder.New()
	v, err := b.Build(s)
	if err != nil {
		log.Printf("Invalid schema, failed to build validator: %s", err)
		return err
	}

	var input interface{}
	if err := v.Validate(input); err != nil {
		log.Printf("Invalid schema: %s", err)
		return err
	}

	return nil
}

func ValidateJSONBytesBySchemaName(jsonBytes *[]byte, schemaName string, repo *repository.Repository) error {
	trySchemaEntry, err := repo.GetSchemaByName(schemaName)
	if err != nil {
		return errors.New("Schema not found")
	}

	schemaString := trySchemaEntry.Data

	schemaLoader := gojsonschema.NewStringLoader(schemaString)
	documentLoader := gojsonschema.NewBytesLoader(*jsonBytes)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return err
	}

	if result.Valid() {
		return nil
	} else {
		//fmt.Printf("The document is not valid. see errors :\n")
		//for _, desc := range result.Errors() {
		//	fmt.Printf("- %s\n", desc)
		//}
		return err
	}
}

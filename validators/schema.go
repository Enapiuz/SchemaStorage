package validators

import (
	"bytes"
	"github.com/lestrrat/go-jsschema"
	"github.com/lestrrat/go-jsval/builder"
	"log"
)

func ValidateBytes(schemaBytes *[]byte) error {
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

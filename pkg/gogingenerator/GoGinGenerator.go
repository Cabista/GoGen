package gogingenerator

import (
	"context"
	"encoding/json"

	"github.com/cabista/gogen/pkg/commongenerator"
	"github.com/getkin/kin-openapi/openapi3"
)

type GoGinGeneratorOptions struct {
	IsGoModule   bool
	GoModuleName string
}

type GoGinGenerator struct {
}

// Generate a client in the given output folder with the given swagger specification and options
func (gen GoGinGenerator) Generate(swagger *openapi3.Swagger, outputLocation string, options string) {
	optionsBytes := []byte(options)
	var goGinGeneratorOptions GoGinGeneratorOptions
	if err := json.Unmarshal(optionsBytes, &goGinGeneratorOptions); err != nil {
		panic(err)
	}

	if err := swagger.Validate(context.Background()); err != nil {
		panic(err)
	}

	swagger.Components

}

// GetType returns the generator type of this generator
func (gen GoGinGenerator) GetType() commongenerator.GeneratorType {
	return commongenerator.GeneratorTypeServer
}

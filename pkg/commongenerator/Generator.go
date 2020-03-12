package commongenerator

import "github.com/getkin/kin-openapi/openapi3"

// GeneratorType Represents a generator type
type GeneratorType uint8

const (
	// GeneratorTypeClient shows this generator generates a client
	GeneratorTypeClient = GeneratorType(iota)
	// GeneratorTypeServer shows this generator generates a server
	GeneratorTypeServer = GeneratorType(iota)
	// GeneratorTypeDocumentation shows this generator generates documentation
	GeneratorTypeDocumentation = GeneratorType(iota)
)

// Generator is a common interface for generators to take a yaml or other file and generate a client
type Generator interface {
	//Generate a client in the given output folder with the given swagger specification and options
	Generate(swagger *openapi3.Swagger, outputLocation string, options string)

	// GetType returns the generator type of this generator
	GetType() GeneratorType
}

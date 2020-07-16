package generator

import (
	"errors"
	"go/ast"
	"log"

	"github.com/google/uuid"
)

// ExampleStruct is a dummy struct to test the program
// Continue Comment
type ExampleStruct struct{}

// EmptyExportedMethod is struct method without param and return value
func (e ExampleStruct) EmptyExportedMethod() {}

// EmptyExportedReferenceMethod is struct method without param and return value
func (e *ExampleStruct) EmptyExportedReferenceMethod() {}

// ParameterMethod is struct method with params and return values
func (e *ExampleStruct) ParameterMethod(param1, param2 string, emptyInterface interface{}) (int, error) {
	return 0, errors.New("Error")
}

// ImportedParameterMethod is struct method with params and return values which are imported
func (e *ExampleStruct) ImportedParameterMethod(u uuid.UUID, s *ast.StructType) (l log.Logger, err error) {
	return
}

// CompositeVariadicMethod is struct method with composite and variadic types
func (e *ExampleStruct) CompositeVariadicMethod(m map[string]uuid.UUID, a [3]uuid.UUID, i <-chan bool, v ...int) (f func(string), err error) {
	return
}

package generator

import (
	"go/types"

	"golang.org/x/tools/go/loader"
)

// Options represents commandline arguments
type Options struct {
	For    string
	As     string
	Output string
}

// Config represents a named type request.
type Config struct {
	StructName  string          `json:"structName,omitempty"`
	PackageName string          `json:"packageName,omitempty"`
	Program     *loader.Program `json:"program,omitempty"`
	Package     *types.Package  `json:"package,omitempty"`
	Object      types.Object    `json:"object,omitempty"`
	Options     Options         `json:"options,omitempty"`
}

// Interface represents an interface signature
type Interface struct {
	Name      string    `json:"name,omitempty"`
	Package   string    `json:"package,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	Functions Functions `json:"functions,omitempty"`
	Imports   Imports   `json:"imports,omitempty"`
}

// Func represents a function signature
type Func struct {
	Name       string  `json:"name,omitempty"`
	Params     []Param `json:"params,omitempty"`
	Res        []Param `json:"res,omitempty"`
	Comment    string  `json:"comment,omitempty"`
	IsVariadic bool    `json:"isVariadic,omitempty"`
}

// Param represents a parameter in a function or method signature
type Param struct {
	Name string `json:"name,omitempty"`
	Type Type   `json:"type,omitempty"`
}

// Type represents a simple representation of a single parameter type
type Type struct {
	Name    string   `json:"name,omitempty"`
	Imports []Import `json:"imports,omitempty"`
}

// Import defines imported package
type Import struct {
	Package string `json:"package,omitempty"`
	Path    string `json:"path,omitempty"`
}

// Functions is a slice of Func
type Functions []Func

func (functions Functions) Len() int           { return len(functions) }
func (functions Functions) Swap(i, j int)      { functions[i], functions[j] = functions[j], functions[i] }
func (functions Functions) Less(i, j int) bool { return functions[i].Name < functions[j].Name }

// Imports is slice of Import
type Imports []Import

func (imports Imports) Len() int           { return len(imports) }
func (imports Imports) Swap(i, j int)      { imports[i], imports[j] = imports[j], imports[i] }
func (imports Imports) Less(i, j int) bool { return imports[i].Path < imports[j].Path }

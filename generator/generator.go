package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sort"
	"strings"
)

// BuildStruct creates an Interface object for provided options
func BuildStruct(options Options) (*Interface, error) {
	config, err := setupConfig(options)
	if err != nil {
		return nil, err
	}

	definition, ok := config.Object.Type().Underlying().(*types.Struct)
	if !ok {
		return nil, fmt.Errorf("Passed type name %q in package %q is not a struct", config.StructName, config.Package.Path())
	}

	funcs, err := structFunctions(config, definition)
	if err != nil {
		return nil, err
	}

	s := &Interface{
		Name:      outputInterfaceName(options.As, config.StructName),
		Comment:   commentText(config, config.Object.Pos()),
		Package:   outputPackageName(options.As, config.Package.Name()),
		Functions: funcs,
	}

	fixupStruct(s)

	return s, nil
}

func setupConfig(options Options) (*Config, error) {
	idx := strings.LastIndex(options.For, ".")
	if idx == -1 || options.For[:idx] == "" || options.For[idx+1:] == "" {
		return nil, fmt.Errorf("--interface (-i) flag should be like path/to/package.type")
	}

	structName := options.For[idx+1:]
	packageName := options.For[:idx]

	program, err := loadProgram(packageName)
	if err != nil {
		return nil, err
	}

	pkg := program.Package(packageName).Pkg
	obj := pkg.Scope().Lookup(structName)
	if obj == nil {
		return nil, fmt.Errorf("Struct %q not found in package %q", structName, packageName)
	}

	return &Config{
		StructName:  structName,
		PackageName: packageName,
		Program:     program,
		Package:     pkg,
		Object:      obj,
		Options:     options,
	}, nil
}

func commentText(config *Config, pos token.Pos) string {
	_, paths, _ := config.Program.PathEnclosingInterval(pos, pos)
	for _, n := range paths {
		switch n := n.(type) {
		case *ast.FuncDecl:
			return commentGroupToString(n.Doc)
		case *ast.GenDecl:
			return commentGroupToString(n.Doc)
		case *ast.Field:
			return commentGroupToString(n.Doc)
		}
	}

	return ""
}

func commentGroupToString(commentGroup *ast.CommentGroup) string {
	s := ""

	if commentGroup == nil {
		return s
	}

	for _, comment := range commentGroup.List {
		s += fmt.Sprintf("%v\n", comment.Text)
	}

	return s
}

func structFunctions(config *Config, s *types.Struct) ([]Func, error) {
	funcs := []Func{}
	info := config.Program.Imported[config.PackageName]
	named, err := structNameType(info.Defs, config.StructName, config.PackageName)
	if err != nil {
		return nil, err
	}

	for i := 0; i < named.NumMethods(); i++ {
		meth := named.Method(i)

		if !meth.Exported() {
			continue
		}

		sig, ok := meth.Type().(*types.Signature)
		if !ok {
			continue
		}

		f := Func{
			Name:       meth.Name(),
			Comment:    commentText(config, meth.Pos()),
			Params:     signatureVariables(sig.Params()),
			Res:        signatureVariables(sig.Results()),
			IsVariadic: sig.Variadic(),
		}

		funcs = append(funcs, f)
	}

	return funcs, nil
}

func structNameType(defs map[*ast.Ident]types.Object, structName, structPackage string) (*types.Named, error) {
	var typ *types.Named

	for _, obj := range defs {
		if obj == nil {
			continue
		}
		if obj.Name() != structName || obj.Pkg().Path() != structPackage {
			continue
		}
		var ok bool
		typ, ok = obj.Type().(*types.Named)
		if ok {
			break
		}
	}
	if typ == nil {
		return nil, fmt.Errorf("No exported methods found for %q (package %q)", structName, structPackage)
	}

	return typ, nil
}

func signatureVariables(tuple *types.Tuple) []Param {
	params := make([]Param, tuple.Len())

	for i := 0; i < tuple.Len(); i++ {
		param := tuple.At(i)
		name := param.Name()
		t := &Type{}
		configureParamType(t, param.Type())

		params[i] = Param{
			Name: name,
			Type: *t,
		}
	}

	return params
}

func configureParamType(t *Type, typ types.Type) {
	switch typ := typ.(type) {
	case *types.Basic:
		configureParamTypeName(t, typ.Name())
	case *types.Pointer:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Named:
		if pkg := typ.Obj().Pkg(); pkg != nil {
			configureParamTypeName(t, fmt.Sprintf("%v.%v", pkg.Name(), typ.Obj().Name()))
			t.Imports = append(t.Imports, Import{Package: pkg.Name(), Path: pkg.Path()})
		}
		configureParamTypeName(t, typ.Obj().Name())
	case *types.Slice:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Signature:
		configureParamTypeName(t, typ.String())
		variables := signatureVariables(typ.Params())
		variables = append(variables, signatureVariables(typ.Results())...)
		for _, v := range variables {
			t.Imports = append(t.Imports, v.Type.Imports...)
		}
	case *types.Array:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Chan:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	case *types.Map:
		configureParamTypeName(t, typ.String())
		configureParamType(t, typ.Elem())
	}
}

func configureParamTypeName(t *Type, name string) {
	if t.Name == "" {
		t.Name = name
	}
}

func fixupStruct(s *Interface) {
	imports := map[string]Import{}

	for fi, f := range s.Functions {
		for pi, p := range f.Params {
			for _, i := range p.Type.Imports {
				imports[i.Path] = i
				s.Functions[fi].Params[pi].Type.Name = strings.ReplaceAll(p.Type.Name, i.Path, i.Package)
			}

			if f.IsVariadic && pi == len(f.Params)-1 {
				s.Functions[fi].Params[pi].Type.Name = strings.Replace(p.Type.Name, "[]", "...", 1)
			}
		}
		for ri, p := range f.Res {
			for _, i := range p.Type.Imports {
				imports[i.Path] = i
				s.Functions[fi].Res[ri].Type.Name = strings.ReplaceAll(p.Type.Name, i.Path, i.Package)
			}
		}
	}

	keys := make([]string, 0, len(imports))
	for k := range imports {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		s.Imports = append(s.Imports, imports[k])
	}
}

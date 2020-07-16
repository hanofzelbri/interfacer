package generator

import (
	"fmt"
	"go/parser"
	"strings"

	"golang.org/x/tools/go/loader"
)

func loadProgram(packageName string) (*loader.Program, error) {
	cfg := &loader.Config{
		AllowErrors:         true,
		ImportPkgs:          map[string]bool{packageName: true},
		TypeCheckFuncBodies: func(string) bool { return false },
		ParserMode:          parser.ParseComments,
	}

	cfg.ImportWithTests(packageName)

	return cfg.Load()
}

func outputPackageName(as string, packageName string) string {
	if i := strings.IndexRune(as, '.'); i != -1 {
		packageName = (as)[:i]
	}

	return packageName
}

func outputInterfaceName(as string, structName string) string {
	interfaceName := fmt.Sprintf("%vInterface", structName)

	if as == "" {
		return interfaceName
	}

	i := strings.IndexRune(as, '.')
	if i == -1 {
		return interfaceName
	}

	return (as)[i+1:]
}

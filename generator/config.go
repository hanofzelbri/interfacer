package generator

import (
	"go/parser"
	"strings"
	"unicode"

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
	interfaceName := as

	if interfaceName == "" {
		interfaceName = string(unicode.ToLower(rune(structName[0]))) + structName[1:]
	}

	if i := strings.IndexRune(as, '.'); i != -1 {
		interfaceName = (as)[i+1:]
	}

	return interfaceName
}

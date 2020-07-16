# Interfacer

Generates an interface for golang struct.

> For detected bugs please contact: marco-engstler@gmx.de

- [Interfacer](#interfacer)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Flags](#flags)
  - [Examples](#examples)
    - [Generate manually](#generate-manually)
    - [Generate by go generate](#generate-by-go-generate)
    - [Example output for _CompositeParamsInterface_ in file interfaces/interfaces_test.go](#example-output-for-compositeparamsinterface-in-file-interfacesinterfaces_testgo)

## Installation

```bash
go get github.com/hanofzelbri/interfacer
```

## Usage

```bash
interfacer [flags]
```

### Flags

```bash
      --as string       Generated output interface definition.
      --for string      Struct definition to generate interface for.
  -h, --help            help for interfacer
  -o, --output string   Output file. If empty StdOut is used
```

## Examples

### Generate manually

```bash
interfacer --for "github.com/minio/minio-go/v6.Client"
interfacer --for "github.com/hanofzelbri/interfacer/interfaces.ExampleStruct" --as "generator.ExampleStructInterface" -o "generator/output.go"
```

### Generate by go generate

```go
//go:generate interfacer --for "github.com/hanofzelbri/interfacer/interfaces.ExampleStruct" -as "generator.ExampleStructInterface" -o "generator/output.go"
```

### Example output for _CompositeParamsInterface_ in file [interfaces/interfaces_test.go](interfaces/interfaces_test.go)

```go

```

## TODOS

- Empty interface as type f.e. interfacer --for github.com/go-pg/pg/v9.DB -o test.go
- Create subfolder if not exists for output
- Add comment for which source Interface and package file was generated
- Sort functions

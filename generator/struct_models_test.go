package generator

var ExampleStructModel = &Interface{
	Name:    "Interface",
	Comment: "// ExampleStruct is a dummy struct to test the program\n",
	Package: "test",
	Functions: []Func{
		{
			Name:       "EmptyExportedMethod",
			Params:     []Param{},
			Res:        []Param{},
			Comment:    "// EmptyExportedMethod is struct method without param and return value\n",
			IsVariadic: false,
		},
		{
			Name:       "EmptyExportedReferenceMethod",
			Params:     []Param{},
			Res:        []Param{},
			Comment:    "// EmptyExportedReferenceMethod is struct method without param and return value\n",
			IsVariadic: false,
		},
		{
			Name: "ParameterMethod",
			Params: []Param{
				{
					Name: "param1",
					Type: Type{
						Name: "string",
					},
				},
				{
					Name: "param2",
					Type: Type{
						Name: "string",
					},
				},
			},
			Res: []Param{
				{
					Name: "",
					Type: Type{
						Name: "int",
					},
				},
				{
					Name: "",
					Type: Type{
						Name: "error",
					},
				},
			},
			Comment:    "// ParameterMethod is struct method with params and return values\n",
			IsVariadic: false,
		},
		{
			Name: "ImportedParameterMethod",
			Params: []Param{
				{
					Name: "u",
					Type: Type{
						Name: "uuid.UUID",
						Imports: []Import{
							{
								Package: "uuid",
								Path:    "github.com/google/uuid",
							},
						},
					},
				},
				{
					Name: "s",
					Type: Type{
						Name: "*ast.StructType",
						Imports: []Import{
							{
								Package: "ast",
								Path:    "go/ast",
							},
						},
					},
				},
			},
			Res: []Param{
				{
					Name: "l",
					Type: Type{
						Name: "log.Logger",
						Imports: []Import{
							{
								Package: "log",
								Path:    "log",
							},
						},
					},
				},
				{
					Name: "err",
					Type: Type{
						Name: "error",
					},
				},
			},
			Comment:    "// ImportedParameterMethod is struct method with params and return values which are imported\n",
			IsVariadic: false,
		},
		{
			Name: "CompositeVariadicMethod",
			Params: []Param{
				{
					Name: "m",
					Type: Type{
						Name: "map[string]uuid.UUID",
						Imports: []Import{
							{
								Package: "uuid",
								Path:    "github.com/google/uuid",
							},
						},
					},
				},
				{
					Name: "a",
					Type: Type{
						Name: "[3]uuid.UUID",
						Imports: []Import{
							{
								Package: "uuid",
								Path:    "github.com/google/uuid",
							},
						},
					},
				},
				{
					Name: "i",
					Type: Type{
						Name: "<-chan bool",
					},
				},
				{
					Name: "v",
					Type: Type{
						Name: "...int",
					},
				},
			},
			Res: []Param{
				{
					Name: "f",
					Type: Type{
						Name: "func(string)",
					},
				},
				{
					Name: "err",
					Type: Type{
						Name: "error",
					},
				},
			},
			Comment:    "// CompositeVariadicMethod is struct method with composite and variadic types\n",
			IsVariadic: true,
		},
	},
	Imports: []Import{
		{
			Package: "uuid",
			Path:    "github.com/google/uuid",
		},
		{
			Package: "ast",
			Path:    "go/ast",
		},
		{
			Package: "log",
			Path:    "log",
		},
	},
}

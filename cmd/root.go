package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hanofzelbri/interfacer/generator"

	"github.com/spf13/cobra"
)

var (
	options = generator.Options{}
)

var rootCmd = &cobra.Command{
	Use:   "interfacer",
	Short: "Generates an interface for golang structs",
	Long: `This golang generator can be used to generate an interface
for a provided struct.

Either use it directly as binary or add it as comment for go:generate --> see examples

For detected bugs please contact: marco-engstler@gmx.de`,
	Example: `interfacer --for "io.Reader" --as "pkg.Structname" -o "file.go"
interfacer --for "github.com/hanofzelbri/interfacer/generator.ExampleStruct" -o "file.go"

//go:generate interfacer --for "github.com/hanofzelbri/interfacer/generator.ExampleStruct" -o "struct.go"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := generator.BuildInterface(options)
		if err != nil {
			return err
		}

		template, err := generator.InterfaceTemplate(s)
		if err != nil {
			return fmt.Errorf("%v\n\nerr: %v", string(template), err)
		}

		if err := os.MkdirAll(filepath.Dir(options.Output), os.ModePerm); err != nil {
			return err
		}

		f := os.Stdout
		if options.Output != "" {
			f, err = os.OpenFile(options.Output, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				return err
			}
		}

		_, err = f.Write(template)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}

		fmt.Printf("Successfully wrote interface to %v\n", options.Output)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.For, "for", "", "Struct definition to generate interface for.")
	rootCmd.MarkPersistentFlagRequired("for")

	rootCmd.PersistentFlags().StringVar(&options.As, "as", "", "Generated output interface definition.")
	rootCmd.PersistentFlags().StringVarP(&options.Output, "output", "o", "", "Output file. If empty StdOut is used")
}

package generator

import (
	"reflect"
	"testing"
)

func TestBuildInterface(t *testing.T) {
	o := Options{
		As:     "test.Interface",
		Output: "",
	}

	tests := []struct {
		name    string
		options func() Options
		want    *Interface
		wantErr bool
	}{
		{
			name: "github.com/hanofzelbri/interfacer/generator.ExampleStruct",
			options: func() Options {
				o.For = "github.com/hanofzelbri/interfacer/generator.ExampleStruct"
				return o
			},
			want:    ExampleStructModel,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildInterface(tt.options())
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

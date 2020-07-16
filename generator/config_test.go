package generator

import "testing"

func Test_outputInterfaceName(t *testing.T) {
	type args struct {
		as         string
		structName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty as",
			args: args{
				as:         "",
				structName: "StructName",
			},
			want: "StructNameInterface",
		},
		{
			name: "Without Package as",
			args: args{
				as:         "InterfaceName",
				structName: "StructName",
			},
			want: "StructNameInterface",
		},
		{
			name: "Without Package as",
			args: args{
				as:         "InterfaceName",
				structName: "StructName",
			},
			want: "StructNameInterface",
		},
		{
			name: "Correct as",
			args: args{
				as:         "test.InterfaceName",
				structName: "StructName",
			},
			want: "InterfaceName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outputInterfaceName(tt.args.as, tt.args.structName); got != tt.want {
				t.Errorf("outputInterfaceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

package factory

import (
	"reflect"
	"testing"
)

func TestCarFactory_New(t *testing.T) {
	carFactory := CarFactory{}

	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    CarFactory
		args args
		want Car
	}{
		{
			name: "CarFactory create BMW",
			c:    carFactory,
			args: args{name: "bmw"},
			want: &BMW{},
		},
		{
			name: "CarFactory create Cadillac",
			c:    carFactory,
			args: args{name: "cadillac"},
			want: &Cadillac{},
		},
		{
			name: "CarFactory create Geely",
			c:    carFactory,
			args: args{name: "geely"},
			want: &Geely{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CarFactory{}
			if got := c.New(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarFactory.New() = %v, want %v", got.Drive(), tt.want.Drive())
			}
		})
	}
}

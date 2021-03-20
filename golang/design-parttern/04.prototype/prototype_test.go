package prototype

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPerson_Clone(t *testing.T) {
	person := &Person{Name: "zhangsan", Age: 18}
	type fields struct {
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Person
	}{
		{
			name:   "person clone equals",
			fields: fields{Name: "zhangsan", Age: 18},
			want:   person.Clone(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := p.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Person.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_Clone2(t *testing.T) {
	person := Person{Name: "zhangsan", Age: 18, Job: &Job{Company: "jD"}}
	person2 := person.Clone()
	person.Name = "lisi"
	person.Job.Company = "Tencent"

	fmt.Printf("%p, %+v, job=%s\n", &person, person, person.Job.Company)
	fmt.Printf("%p, %+v, job=%s\n", person2, *person2, person2.Job.Company)
	// fmt.Printf("%p, %+v, job=%s\n", &person2, person2, person2.Job.Company)
}

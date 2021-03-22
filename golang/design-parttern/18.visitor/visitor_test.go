package visitor

import "testing"

func TestMyExample_Print(t *testing.T) {
	ele := Element{}
	// env := "production"
	env := "testing"
	ele.Accept(&ProductionVisitor{Env: env})
	ele.Accept(&TestingVisitor{Env: env})

	example := &ExampleLog{Element: ele}
	example.Print()
}

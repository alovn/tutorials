package builder

import (
	"fmt"
	"testing"
)

func TestMacbookBuilder_Build(t *testing.T) {
	var builder ComputerBuilder = &MacbookBuilder{}
	builder.SetCPU("intel").SetGPU("nvidia").SetBoard("asus")

	var computer Computer = builder.Build()
	fmt.Printf("%+v\n", computer)
}

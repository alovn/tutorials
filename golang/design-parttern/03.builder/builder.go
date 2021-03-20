package builder

type Computer struct {
	CPU   string
	GPU   string
	Board string
}

type ComputerBuilder interface {
	SetCPU(name string) ComputerBuilder
	SetGPU(name string) ComputerBuilder
	SetBoard(name string) ComputerBuilder
	Build() Computer
}

type MacbookBuilder struct {
	cpu, gpu, board string
}

func (m *MacbookBuilder) SetCPU(name string) ComputerBuilder {
	m.cpu = name
	return m
}
func (m *MacbookBuilder) SetGPU(name string) ComputerBuilder {
	m.gpu = name
	return m
}
func (m *MacbookBuilder) SetBoard(name string) ComputerBuilder {
	m.board = name
	return m
}
func (m *MacbookBuilder) Build() Computer {
	return Computer{CPU: m.cpu, GPU: m.gpu, Board: m.board}
}

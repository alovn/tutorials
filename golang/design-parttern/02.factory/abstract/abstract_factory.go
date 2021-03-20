package factory

import (
	"fmt"
)

//计算机抽象工厂
type ComputerFactory interface {
	CreateCPU() CPU
	CreateGPU() GPU
}

//CPU抽象
type CPU interface {
	Compute()
}

//GPU抽象
type GPU interface {
	Display()
}

//Intel CPU
type IntelCPU struct{}

func (i IntelCPU) Compute() {
	fmt.Println("Intel CPU Compute")
}

//Intel GPU
type IntelGPU struct{}

func (i IntelGPU) Display() {
	fmt.Println("Intel GPU Display")
}

//AMD CPU
type AMDCPU struct{}

func (a AMDCPU) Compute() {
	fmt.Println("AMD CPU Compute")
}

//NVidia GPU
type NVidiaGPU struct{}

func (n NVidiaGPU) Display() {
	fmt.Println("NVidia GPU Display")
}

//苹果工厂
type AppleFactory struct{}

func (a AppleFactory) CreateCPU() CPU {
	return &IntelCPU{}
}

func (a AppleFactory) CreateGPU() GPU {
	return &NVidiaGPU{}
}

//华为工厂
type HuaweiFactory struct{}

func (m HuaweiFactory) CreateCPU() CPU {
	return &AMDCPU{}
}
func (m HuaweiFactory) CreateGPU() GPU {
	return &NVidiaGPU{}
}

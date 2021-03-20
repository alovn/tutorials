package factory

import (
	"testing"
)

func TestFactory_Create(t *testing.T) {
	var huawei HuaweiFactory
	huawei.CreateCPU().Compute()
	huawei.CreateGPU().Display()

	var apple AppleFactory
	apple.CreateCPU().Compute()
	apple.CreateGPU().Display()
}

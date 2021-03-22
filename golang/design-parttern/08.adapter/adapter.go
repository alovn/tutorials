package adapter

import "fmt"

//美国供电口
type AmericaPower struct {
}

func (a AmericaPower) Power() {
	fmt.Println("美国电源供电")
}

type ChinaPlug struct{}

func (c ChinaPlug) Charge() {
	fmt.Println("中式插口充电")
}

type PowerAdapter struct {
	power *AmericaPower
}

func (a *PowerAdapter) SetPower(i *AmericaPower) {
	a.power = i
}

func (a PowerAdapter) Charge(i *ChinaPlug) {
	a.power.Power()
	fmt.Println("电压转换...")
	i.Charge()
}

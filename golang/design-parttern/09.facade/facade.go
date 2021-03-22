package facade

import "fmt"

type Light struct{}

func (l Light) Open() {
	fmt.Println("打开灯")
}

type TV struct{}

func (tv TV) Open() {
	fmt.Println("打开电视机")
}

type AirConditioning struct{}

func (a AirConditioning) Open() {
	fmt.Println("打开空调")
}

type WaterHeater struct{}

func (w WaterHeater) Open() {
	fmt.Println("打开热水器")
}

type Facade struct{}

func (f Facade) Open() {
	Light{}.Open()
	TV{}.Open()
	AirConditioning{}.Open()
	WaterHeater{}.Open()
}

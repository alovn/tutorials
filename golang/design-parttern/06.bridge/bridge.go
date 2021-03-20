package bridge

import "fmt"

type Car interface {
	UseEngine(Engine)
	Start()
}

//发动机
type Engine interface {
	Start()
}

//Volvo发动机
type VolvoEngine struct{}

func (v VolvoEngine) Start() {
	fmt.Println("Start Volvo Engine")
}

//吉利汽车
type GeelyCar struct {
	Engine Engine
}

//使用什么发动机
func (g *GeelyCar) UseEngine(engine Engine) {
	g.Engine = engine
}

//汽车启动
func (g GeelyCar) Start() {
	g.Engine.Start()
	fmt.Println("GeelyCar Started")
}

package visitor

import "fmt"

//访问者接口
type IVisitor interface {
	Visite() //访问者的访问方法
}

//具体访问者
type ProductionVisitor struct {
	Env string
}

func (v ProductionVisitor) Visite() {
	if v.Env == "production" {
		fmt.Println("这是生成环境的输出")
	}
}

type TestingVisitor struct {
	Env string
}

func (v TestingVisitor) Visite() {
	if v.Env == "testing" {
		fmt.Println("这是测试环境的输出")
	}
}

//IElement抽象元素，在其中声明一个accept()操作，它以一个抽象访问者作为参数
type IElement interface {
	Accept(visitor IVisitor)
}

//具体元素，它实现了accept()操作，在accept()中调用访问者的访问方法以便完成对一个元素的操作
type Element struct {
	visitors []IVisitor
}

func (ele *Element) Accept(visitor IVisitor) {
	ele.visitors = append(ele.visitors, visitor)
}

//修改打印输出方法
type ExampleLog struct {
	Element
}

func (ex ExampleLog) Print() {
	for _, v := range ex.visitors {
		v.Visite()
	}
}

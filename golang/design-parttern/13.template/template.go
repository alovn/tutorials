package template

import "fmt"

type TaskTemplate struct {
	Before func()
	After  func()
}

func (t TaskTemplate) doTask() {
	fmt.Println("do task template")
}
func (t TaskTemplate) Do() {
	t.Before()
	t.doTask()
	t.After()
}

type MyTaskTemplate struct {
	TaskTemplate
}

func NewMyTaskTemplate() *MyTaskTemplate {
	myTask := &MyTaskTemplate{}
	myTask.TaskTemplate.Before = myTask.Before
	myTask.TaskTemplate.After = myTask.After
	return myTask
}

func (t MyTaskTemplate) Before() {
	fmt.Println("befor work")
}

func (t MyTaskTemplate) After() {
	fmt.Println("after work")
}

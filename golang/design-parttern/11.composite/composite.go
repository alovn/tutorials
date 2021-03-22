package composite

import "fmt"

type Widget interface {
	Draw()
}

type BaseWidget struct {
	Type       string
	Text       string
	subWidgets []Widget
}

//添加子组件
func (w *BaseWidget) Add(widget Widget) {
	w.subWidgets = append(w.subWidgets, widget)
}

func (w BaseWidget) Draw() {
	fmt.Printf("Draw type=%s, text=%s\n", w.Type, w.Text)
	for _, widget := range w.subWidgets {
		widget.Draw()
	}
}

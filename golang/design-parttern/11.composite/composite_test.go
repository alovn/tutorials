package composite

import "testing"

func TestBaseWidget_Draw(t *testing.T) {
	window := &BaseWidget{Text: "Window", Type: "Form"}
	window.Add(BaseWidget{Text: "Logo", Type: "Picture"})
	window.Add(BaseWidget{Text: "Login", Type: "Button"})
	window.Add(BaseWidget{Text: "Cancel", Type: "Button"})
	window.Draw()
}

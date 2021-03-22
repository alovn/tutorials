package memento

import "fmt"

//备忘录快照，记录某时刻的状态，这里可以记录更多的状态
type Memento struct {
	state string
}

func (m *Memento) SetState(state string) {
	m.state = state
}
func (m *Memento) GetState() string {
	return m.state
}

type TextEditor struct {
	mementos map[string]Memento
	str      string
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		mementos: make(map[string]Memento),
	}
}

func (editor *TextEditor) Append(s string) {
	editor.str += s
}
func (editor *TextEditor) Content() string {
	return editor.str
}

func (editor *TextEditor) Delete() {
	if len(editor.str) > 0 {
		editor.str = editor.str[0 : len(editor.str)-1]
	}
}

//获取某状态当时的值
func (editor *TextEditor) GetState(state string) string {
	if memento, ok := editor.mementos[state]; ok {
		fmt.Printf("获取到状态为%s时的值：%s\n", state, memento.GetState())
		return memento.state
	}
	return ""
}

//保存当前状态时的值
func (editor *TextEditor) SaveState(state string) {
	memento := Memento{}
	memento.SetState(editor.str)
	editor.mementos[state] = memento
	fmt.Printf("保存状态为%s时的值：%s\n", state, memento.GetState())
}

//绘图到某状态时的值
func (editor *TextEditor) BackState(state string) {
	if memento, ok := editor.mementos[state]; ok {
		fmt.Printf("回退到状态为%s时的值：%s\n", state, memento.GetState())
		editor.str = memento.state
	}
}

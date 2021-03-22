package memento

import (
	"fmt"
	"testing"
)

func TestTextEditor_SetState(t *testing.T) {
	editor := NewTextEditor()
	editor.SaveState("state#0") //保存状态

	editor.Append("1")
	editor.SaveState("state#1") //保存状态

	editor.Append("2")
	editor.SaveState("state#2") //保存状态

	editor.Append("3")
	editor.SaveState("state#3") //保存状态

	_ = editor.GetState("state#2") //获取当时状态

	_ = editor.GetState("state#1") //获取当时状态

	editor.BackState("state#1") //回退到当时状态
	fmt.Println("当前值为:", editor.Content())
}

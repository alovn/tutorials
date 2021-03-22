package chain

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handle(content string)
}

type LogHandler struct {
	NextHandler Handler
}

//日志记录
func (h LogHandler) Handle(content string) {
	fmt.Println("[执行日志记录]")
	fmt.Println(content)
	if h.NextHandler != nil {
		h.NextHandler.Handle(content)
	}
}

//广告过滤
type AdFilterHandler struct {
	NextHandler Handler
}

func (h AdFilterHandler) Handle(content string) {
	fmt.Println("[执行广告过滤]")
	newContent := strings.ReplaceAll(content, "广告", "**")
	if h.NextHandler != nil {
		h.NextHandler.Handle(newContent)
	}
}

//过滤脚本字符
type ScriptFilterHandler struct {
	NextHandler Handler
}

func (h ScriptFilterHandler) Handle(content string) {
	fmt.Println("[执行脚本过滤]")
	newContent := strings.ReplaceAll(content, "script", "")
	if h.NextHandler != nil {
		h.NextHandler.Handle(newContent)
	}
}

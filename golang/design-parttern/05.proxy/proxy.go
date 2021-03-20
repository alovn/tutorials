package proxy

import (
	"fmt"
)

type User struct {
	Name string
}

func (u User) Login() {
	fmt.Printf("%s 正在登录\n", u.Name)
}

type Handler interface {
	HandleFunc(pattern string, handler func())
}

type LoginHandler struct{}

func (l LoginHandler) HandleFunc(pattern string, fn func()) {
	fmt.Printf("调用登录: %s\n", pattern)
	fn()
	fmt.Printf("完成登录: %s\n", pattern)
}

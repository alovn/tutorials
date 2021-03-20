package proxy

import (
	"testing"
)

func TestUser_Login(t *testing.T) {
	var loginHander LoginHandler
	loginHander.HandleFunc("/user/login", func() {
		user := &User{Name: "zhangsan"}
		user.Login()
	})
}

package state

import "fmt"

type State int

const (
	UNKOWN State = iota
	ONLINE
	BUSY
)

type IState interface {
	ReceiveMessage(msg string)
}
type OnlineState struct {
}

func (s OnlineState) ReceiveMessage(msg string) {
	fmt.Println("收到消息并提醒：", msg)
}

type BusyState struct {
}

func (s BusyState) ReceiveMessage(msg string) {
	fmt.Println("回复忙碌中")
}

type StateContext struct {
	state IState
}

func (ctx *StateContext) SetState(s IState) {
	ctx.state = s
}

func (ctx StateContext) ReceiveMessage(msg string) {
	if msg == "online" {
		ctx.SetState(OnlineState{})
	}
	ctx.state.ReceiveMessage(msg)
}

package state

import "testing"

func TestStateContext_ReceiveMessage(t *testing.T) {
	onlineState := &OnlineState{}
	ctx := &StateContext{}
	ctx.SetState(onlineState)
	ctx.ReceiveMessage("hello")

	ctx.SetState(&BusyState{})
	ctx.ReceiveMessage("hello")
	ctx.ReceiveMessage("online")
}

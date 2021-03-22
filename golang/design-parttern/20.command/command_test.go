package command

import "testing"

func TestTextEditor_Append(t *testing.T) {
	waiter := &Waiter{}
	cmd := NewPourWaterCommand(waiter)
	cmd.Execute()

	cmd = NewCookCommand(&Chef{})
	cmd.Execute()

	cmd = NewCheckoutCommand(waiter)
	cmd.Execute()
}

package adapter

import "testing"

func TestPowerAdapter_Charge(t *testing.T) {
	adapter := &PowerAdapter{}
	adapter.SetPower(&AmericaPower{})
	adapter.Charge(&ChinaPlug{})
}

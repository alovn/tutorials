package bridge

import "testing"

func TestGeelyCar_Start(t *testing.T) {
	var volvoEngine Engine = &VolvoEngine{}
	var geelyCar Car = &GeelyCar{}
	geelyCar.UseEngine(volvoEngine)
	geelyCar.Start()
}

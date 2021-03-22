package mediator

import (
	"fmt"
	"testing"
)

func TestHouseMediator_Register(t *testing.T) {
	mediator := &HouseMediator{}
	mediator.Register(NewSmallHouseOwner())
	mediator.Register(NewBigHouseOwner())

	person := NewPerson(mediator)
	person.FindHouse("大房子")
	fmt.Println("")

	person.FindHouse("大房子")
}

package interpreter

import (
	"fmt"
	"testing"
)

func TestCaculator_Caculate(t *testing.T) {
	calc := NewCalculator("1 + 10")
	result := calc.Calculate()
	fmt.Println(result)
}

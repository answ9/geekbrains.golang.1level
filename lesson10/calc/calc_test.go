package calc_test

import (
	"lesson10/calc"
	"testing"
)

// все расписывать не стал, все остальное примерно аналогично
func TestOperationPlus(t *testing.T) {
	if res := calc.Calc("+", 2, 2); res != 4 {
		t.Errorf("Calc operation PLUS not work. 2+2 should be 4, but result %f", res)
	}
}


func TestOperationMod(t *testing.T) {
	if res := calc.Calc("%", 35, 3); res != 2 {
		t.Errorf("Calc operation MOD not work. 35 mod 3 should be 2, but result %f", res)
	}
}

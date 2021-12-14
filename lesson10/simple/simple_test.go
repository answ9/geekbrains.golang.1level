package simple_test

import (
	"lesson10/simple"
	"testing"
)

func TestSimpleCalc100(t *testing.T) {
	referenceList := []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	calcList := simple.SimpleCalc(100)

	for i, n := range referenceList {
		if referenceList[i] != calcList[i] {
			t.Errorf("Simple Calc wrong in %d position, should be %d, but received %d", i, n, calcList[i])
		}
	}

}

package math

import (
	"testing"
)

func TestPow(t *testing.T) {
	pow := Pow(2.0, 2)
	if pow != 4.000 {
		t.Errorf("result=%v", pow)
	}
}

func TestFastPow(t *testing.T) {
	pow := FastPow(2.0, 2)
	if pow != 4.000 {
		t.Errorf("result=%v", pow)
	}
}

func TestMooresVoting(t *testing.T) {
	input := []int{1, 1, 1, 3, 2, 2, 4, 2, 2}
	expected, resp := 2, MooresVoting(input)

	if expected != resp {
		t.Errorf("failed")
	}
}

package dynamic_programming

import (
	"testing"
)

func TestKandaneSum(t *testing.T) {
	input := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	expected, resp := 7, KandanesMaxSubarraySum(input)

	if resp != expected {
		t.Errorf("failed: expected %v, got %v", expected, resp)
	}
}

func TestKandaneDifference(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	expected, resp := 5, KandanesMaxDifference(input)

	if resp != expected {
		t.Errorf("failed: expected %v, got %v", expected, resp)
	}
}

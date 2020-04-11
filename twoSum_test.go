package main

import "testing"

func TestTwoSum(t *testing.T) {
	res := TwoSum([]int{1, 2, 3}, 4)
	if res != [2]int{0, 2} {
		t.Error("Must be equal [0,2]")
	}
}

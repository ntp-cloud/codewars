package main

import "testing"

func TestMultiple3And5(t *testing.T) {
	total := Multiple3And5(10)
	if total != 23 {
		t.Errorf("Expected 23!")
	}
}

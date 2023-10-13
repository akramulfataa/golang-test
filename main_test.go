package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Logf("Expected %d, got %d", result, 3)
		t.Fail()
	}
}

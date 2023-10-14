package main

import "testing"

func TestAdd(t *testing.T) {

	testCases := []struct {
		a, b          int
		expectedHasil int
	}{
		{a: 1, b: 2, expectedHasil: 3},
		{a: -1, b: 2, expectedHasil: 1},
		{a: 3, b: 2, expectedHasil: 5},
		{a: 4, b: 10, expectedHasil: 14},
	}

	for _, test := range testCases {
		hasil := Add(test.a, test.b)
		if hasil != test.expectedHasil {
			t.Logf("Expected %d, got %d", test.expectedHasil, hasil)
			t.Fail()
		}
	}

	// result := Add(1, 2)
	// if result != 3 {
	// 	t.Logf("Expected %d, got %d", result, 3)
	// 	t.Fail()
	// }
}

func TestWithHierarchical(t *testing.T) {

	t.Run("a=positive", func(t *testing.T) {
		a := 10
		t.Run("b=positive", func(t *testing.T) {
			result := Add(a, 5)
			if result != 15 {
				t.Logf("Expected %d, got %d", result, 15)
				t.Fail()
			}
		})

		t.Run("b=negative", func(t *testing.T) {
			result := Add(a, -5)
			if result != 5 {
				t.Logf("Expected %d, got %d", result, 5)
				t.Fail()
			}
		})
	})
}

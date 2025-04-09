package main

import "testing"

func TestFactorial(t *testing.T) {
	testCase := []struct {
		name     string
		num      int
		expected int
	}{
		{"Case 2", 2, 2},
		{"Case 5", 5, 120},
		{"Case 0", -1, 0},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.num)
			expectedResult := tc.expected
			if result != expectedResult {
				t.Errorf("Factorial(%d) = %d is wring, correct is %d",
					tc.num, result, tc.expected)
			}
		})
	}
}

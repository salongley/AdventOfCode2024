package main

import "testing"

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
// In the example above, the reports can be found safe or unsafe by checking those rules:

// 7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
// 1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
// 9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
// 1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
// 8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
// 1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

func Test_safe_report(t *testing.T) {

	tests := []struct {
		name           string
		expectedResult bool
		data           []int
	}{
		{"Safe_increasing", true, []int{0, 1, 2, 4, 7, 8, 11}},
		{"Not_Safe_no_change", false, []int{5, 5, 5, 5, 5, 5, 5, 5}},
		{"Safe_decreasing", true, []int{11, 10, 8, 5, 3}},
		{"NotSafe_decreasing", false, []int{11, 10, 8, 6, 1}},
		{"NotSafe_increasing", false, []int{0, 1, 2, 4, 8, 11}},
		{"NotSafe_increasing_and_decreasing", false, []int{0, 1, 2, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		result := safe_report(test.data)
		if result != test.expectedResult {
			t.Fatalf("Test: %v : got %v, expected %v", test.name, result, test.expectedResult)
		}
	}
}

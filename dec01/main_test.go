package main

import (
	"testing"
)

func Test_calculate_distance(t *testing.T) {

	expected := 11

	set1, set2, err := readDataFile("test.txt")
	check(err)

	result, err := calculate_distance(set1, set2)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}

}

func Test_similarity_score(t *testing.T) {
	expected := 31

	set1, set2, err := readDataFile("test.txt")
	check(err)

	result := similarity_score(set1, set2)
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

type testcases struct {
	expected int
	search   int
	list     []int
}

func Test_count_occurences(t *testing.T) {
	testCases := make(map[string]testcases)
	testCases["NoneFound"] = testcases{expected: 0, search: 5, list: []int{0, 0, 0, 0, 0}}
	testCases["AllMatch"] = testcases{expected: 5, search: 5, list: []int{5, 5, 5, 5, 5}}
	testCases["OneFound"] = testcases{expected: 1, search: 5, list: []int{0, 5, 0, 0, 0}}
	var result int
	for key, value := range testCases {
		result = count_occurences(value.search, value.list)
		if result != value.expected {
			t.Fatalf("Test %v, Expected %v, got %v", key, value.expected, result)
		}
	}
}

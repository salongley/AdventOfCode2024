package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	set1, set2, err := readDataFile("input.txt")
	check(err)
	result, err := calculate_distance(set1, set2)
	check(err)
	fmt.Printf("Calculte Distnace %v\n", result)
	similarity := similarity_score(set1, set2)
	fmt.Printf("Calculte Similarity %v \n", similarity)
}

func calculate_distance(set1 []int, set2 []int) (int, error) {
	if len(set1) != len(set2) {
		return 0, fmt.Errorf("sets not same length")
	}
	sort.Ints(set1)
	sort.Ints(set2)
	var total int = 0
	for count := range len(set1) {
		dist_dff := set1[count] - set2[count]
		distance := math.Abs(float64(dist_dff))
		total = total + int(distance)
	}
	return total, nil
}

func readDataFile(filename string) ([]int, []int, error) {
	var set1 []int
	var set2 []int
	file, err := os.Open(filename)
	if err != nil {
		return set1, set2, fmt.Errorf("Could not open file %v", err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		values := strings.Split(lineText, "   ")
		val1, _ := strconv.Atoi(values[0])

		val2, _ := strconv.Atoi(values[1])

		set1 = append(set1, val1)
		set2 = append(set2, val2)
	}
	return set1, set2, nil
}

func similarity_score(set1 []int, set2 []int) int {

	total := 0

	for _, item := range set1 {
		if slices.Contains(set2, item) {
			countOfOccurrences := count_occurences(item, set2)
			total = total + (item * countOfOccurrences)

		}
	}
	return total
}

func count_occurences(item int, set []int) int {
	count := 0
	for _, value := range set {
		if value == item {
			count++
		}
	}
	return count
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	corruptDate, err := readDataFile("input.txt")
	if err != nil {
		panic(err)
	}
	matches := find_matches(corruptDate)
	total := 0
	for _, match := range matches {
		result, err := calculateMul(match)
		if err != nil {
			panic(err)
		}
		total = total + result
	}
	fmt.Printf("total %v", total)
}

func find_matches(text string) []string {

	r, _ := regexp.Compile(`mul(\([0-9]+,[0-9]+)\)`)

	matches := r.FindAllString(text, -1)
	return matches
}

func calculateMul(mulcalc string) (int, error) {
	text := strings.Replace(mulcalc, "mul(", "", -1)
	text = strings.Replace(text, ")", "", -1)
	values := strings.Split(text, ",")
	val1, err1 := strconv.Atoi(values[0])
	val2, err2 := strconv.Atoi(values[1])

	if err1 != nil || err2 != nil {

		return 0, fmt.Errorf("%v, %v", err1, err2)
	}

	return val1 * val2, nil
}

func readDataFile(filename string) (string, error) {
	b, err := os.ReadFile(filename) // just pass the file name
	if err != nil {
		return "", fmt.Errorf("could not read file: %v", err)
	}

	str := string(b) // convert content to a 'string'

	return str, nil
}

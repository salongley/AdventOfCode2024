package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
// In the example above, the reports can be found safe or unsafe by checking those rules:

// 7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
// 1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
// 9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
// 1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
// 8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
// 1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
//

type report struct {
	data []int
}

func main() {
	safeReportCount := 0
	reportData, err := readDataFile("reports.txt")
	if err != nil {
		panic(err)
	}

	for _, item := range reportData {
		if safe_report(item.data) {
			safeReportCount++
		}
	}
	fmt.Println(safeReportCount)
}

func safe_report(report []int) bool {

	// Determine initial direction
	isIncreasing := report[0] < report[len(report)-1]

	for i := 1; i < len(report); i++ {
		delta := math.Abs(float64(report[i]) - float64(report[i-1]))
		if delta == 0 {
			return false
		}
		if isIncreasing {
			// For increasing, each element should be >= previous

			if report[i] < report[i-1] || delta > 3.0 {
				return false
			}
		} else {
			// For decreasing, each element should be <= previous
			if report[i] > report[i-1] || delta > 3.0 {
				return false
			}
		}

	}
	return true
}

func readDataFile(filename string) (map[int]report, error) {
	reports := make(map[int]report)
	var lineCount int
	file, err := os.Open(filename)
	if err != nil {
		return reports, fmt.Errorf("could not open file %v", err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineCount++
		lineText := fileScanner.Text()

		data := strings.Split(lineText, " ")
		var reportData []int
		for _, value := range data {
			dataInt, err := strconv.Atoi(value)
			if err != nil {
				return reports, fmt.Errorf("could not convert string %v", value)
			}
			reportData = append(reportData, dataInt)
		}
		reports[lineCount] = report{data: reportData}

	}
	return reports, nil
}

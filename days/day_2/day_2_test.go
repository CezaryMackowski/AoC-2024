package day_2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getInput(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	var reports [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		tempSlice := make([]int, len(numbers))
		for i := 0; i <= len(numbers)-1; i++ {
			number, err2 := strconv.Atoi(numbers[i])
			if err2 != nil {
				panic("Cannot convert to numbers")
			}
			tempSlice[i] = number
		}
		reports = append(reports, tempSlice)
	}

	return reports
}

func TestDay2(t *testing.T) {
	reports := getInput("../inputs/day_2.txt")

	var tests = []struct {
		reports        [][]int
		expectedResult int
	}{
		{[][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}, 2},
		{reports, 564},
	}

	for _, tt := range tests {
		t.Run("Day 2", func(t *testing.T) {
			result := day2(tt.reports)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay2Extra(t *testing.T) {
	reports := getInput("../inputs/day_2_extra.txt")

	var tests = []struct {
		reports        [][]int
		expectedResult int
	}{
		{[][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}, 4},
		{reports, 604},
	}

	for _, tt := range tests {
		t.Run("Day 2 (Extra)", func(t *testing.T) {
			result := day2_extra(tt.reports)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

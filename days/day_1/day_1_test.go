package day_1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getInput(filename string) ([]int, []int) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil
	}
	defer f.Close()

	var firstColumn []int
	var secondColumn []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		if len(numbers) == 2 {
			first, err1 := strconv.Atoi(numbers[0])
			second, err2 := strconv.Atoi(numbers[1])

			if err1 != nil || err2 != nil {
				panic("Cannot convert numbers")
			}

			firstColumn = append(firstColumn, first)
			secondColumn = append(secondColumn, second)
		}
	}

	return firstColumn, secondColumn
}

func TestDay1(t *testing.T) {
	left, right := getInput("../inputs/day_1.txt")

	var tests = []struct {
		left           []int
		right          []int
		expectedResult int
	}{
		{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 11},
		{left, right, 1879048},
	}

	for _, tt := range tests {
		t.Run("Day 1", func(t *testing.T) {
			result := day1(tt.left, tt.right)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay1Extra(t *testing.T) {
	left, right := getInput("../inputs/day_1_extra.txt")

	var tests = []struct {
		left           []int
		right          []int
		expectedResult int
	}{
		{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 31},
		{left, right, 21024792},
	}

	for _, tt := range tests {
		t.Run("Day 1 (Extra)", func(t *testing.T) {
			result := day1_extra(tt.left, tt.right)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

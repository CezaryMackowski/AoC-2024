package day_3

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func getInput(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()

	var builder strings.Builder

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	return builder.String()
}

func TestDay3(t *testing.T) {
	input := getInput("../inputs/day_3.txt")

	var tests = []struct {
		input          string
		expectedResult int
	}{
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
		{input, 174103751},
	}

	for _, tt := range tests {
		t.Run("Day 3", func(t *testing.T) {
			result := day3(tt.input)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay3Extra(t *testing.T) {
	input := getInput("../inputs/day_3_extra.txt")

	var tests = []struct {
		input          string
		expectedResult int
	}{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
		{input, 100411201},
	}

	for _, tt := range tests {
		t.Run("Day 3 (Extra)", func(t *testing.T) {
			result := day3_extra(tt.input)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

package day_3

import (
	"regexp"
	"strconv"
)

func day3(input string) int {
	re := regexp.MustCompile(`(mul\((\d{1,3})\,(\d{1,3})\))`)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0
	for _, match := range matches {
		num1, err1 := strconv.Atoi(match[2])
		num2, err2 := strconv.Atoi(match[3])
		if err1 != nil || err2 != nil {
			panic("Cannot convert numbers")
		}
		result += num1 * num2
	}

	return result
}

func day3_extra(input string) int {
	re := regexp.MustCompile(`((mul\((\d{1,3})\,(\d{1,3})\))|(do\(\))|(don\'t\(\)))`)
	matches := re.FindAllStringSubmatch(input, -1)

	result := 0
	operation := "do()"
	for _, match := range matches {
		if match[0] == "do()" || match[0] == "don't()" {
			operation = match[0]
			continue
		}

		if operation == "do()" {
			num1, err1 := strconv.Atoi(match[3])
			num2, err2 := strconv.Atoi(match[4])
			if err1 != nil || err2 != nil {
				panic("Cannot convert numbers")
			}
			result += num1 * num2
		}
	}

	return result
}

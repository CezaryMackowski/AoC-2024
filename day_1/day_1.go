package day_1

import "math"

func getSmallestNumberInSlice(s []int) (int, int) {
	smallest := math.MaxInt
	index := 0

	for i := 0; i < len(s); i++ {
		if s[i] < smallest {
			smallest = s[i]
			index = i
		}
	}

	return smallest, index
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func day1(left []int, right []int) int {
	lengthLeft := len(left)
	lengthRight := len(right)
	var finalNumbers []int
	var finalResult int

	if lengthLeft != lengthRight {
		panic("Different lengths")
	}

	for i := 0; i < lengthLeft; i++ {
		smallestLeft, smallestLeftIndex := getSmallestNumberInSlice(left)
		smallestRight, smallestRightIndex := getSmallestNumberInSlice(right)

		finalNumbers = append(finalNumbers, int(math.Abs(float64(smallestLeft-smallestRight))))

		left = remove(left, smallestLeftIndex)
		right = remove(right, smallestRightIndex)
	}

	for i := 0; i < len(finalNumbers); i++ {
		finalResult += finalNumbers[i]
	}
	return finalResult
}

func countNumberInSlice(n int, s []int) int {
	var count int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == n {
			count++
		}
	}
	return count
}

func day1_extra(left []int, right []int) int {
	var finalResult int

	for i := 0; i < len(left); i++ {
		finalResult += left[i] * countNumberInSlice(left[i], right)
	}

	return finalResult
}

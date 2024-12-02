package day_1

import (
	"github.com/CezikLikeWhat/AoC-2024/utils"
	"math"
)

func day1(left []int, right []int) int {
	lengthLeft := len(left)
	lengthRight := len(right)
	var finalNumbers []int
	var finalResult int

	if lengthLeft != lengthRight {
		panic("Different lengths")
	}

	for i := 0; i < lengthLeft; i++ {
		smallestLeft, smallestLeftIndex := utils.GetSmallestElementInSlice(left)
		smallestRight, smallestRightIndex := utils.GetSmallestElementInSlice(right)

		finalNumbers = append(finalNumbers, int(math.Abs(float64(smallestLeft-smallestRight))))

		left = utils.RemoveElementFromSlice(left, smallestLeftIndex)
		right = utils.RemoveElementFromSlice(right, smallestRightIndex)
	}

	for i := 0; i < len(finalNumbers); i++ {
		finalResult += finalNumbers[i]
	}
	return finalResult
}

func day1_extra(left []int, right []int) int {
	var finalResult int

	for i := 0; i < len(left); i++ {
		finalResult += left[i] * utils.CountNumberInSlice(left[i], right)
	}

	return finalResult
}

package day_2

import (
	"math"
)

func day2(reports [][]int) int {
	safeCounter := 0

	for _, report := range reports {
		if isValidReport(report) {
			safeCounter++
		}
	}

	return safeCounter
}

func day2_extra(reports [][]int) int {
	safeCounter := 0

	for _, report := range reports {
		if isValidReport(report) {
			safeCounter++
			continue
		}

		for i := 0; i < len(report); i++ {
			modifiedReport := append([]int{}, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)

			if isValidReport(modifiedReport) {
				safeCounter++
				break
			}
		}
	}

	return safeCounter
}

func isValidReport(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isGrowing, isDecreasing := false, false

	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			isGrowing = true
		} else if report[i] > report[i+1] {
			isDecreasing = true
		}

		if isGrowing && isDecreasing {
			return false
		}

		diff := int(math.Abs(float64(report[i] - report[i+1])))

		if diff > 3 || diff < 1 {
			return false
		}
	}

	return true
}

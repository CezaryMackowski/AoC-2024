package utils

import (
	"golang.org/x/exp/constraints"
)

func GetSmallestElementInSlice[T constraints.Float | constraints.Integer](s []T) (T, int) {
	if len(s) == 0 {
		var zero T
		return zero, -1
	}

	smallest := s[0]
	index := 0

	for i := 0; i < len(s); i++ {
		if s[i] < smallest {
			smallest = s[i]
			index = i
		}
	}

	return smallest, index
}

func CountNumberInSlice[T constraints.Float | constraints.Integer](n T, s []T) int {
	var count = 0
	for i := 0; i < len(s); i++ {
		if s[i] == n {
			count++
		}
	}
	return count
}

func RemoveElementFromSlice[T constraints.Float | constraints.Integer](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

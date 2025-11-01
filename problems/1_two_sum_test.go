package problems

import (
	"slices"
	"testing"
)

func TestExampleOneTest(t *testing.T) {
	exp := []int{0, 1}
	res := twoSum([]int{2, 7, 11, 15}, 9)

	if !slices.Contains(res, exp[0]) || !slices.Contains(res, exp[1]) {
		t.Errorf(`twoSum result did not contain %d and %d!`, exp[0], exp[1])
	}
}

func TestExampleTwoTest(t *testing.T) {
	exp := []int{1, 2}
	res := twoSum([]int{3, 2, 4}, 6)

	if !slices.Contains(res, exp[0]) || !slices.Contains(res, exp[1]) {
		t.Errorf(`twoSum result did not contain %d and %d!`, exp[0], exp[1])
	}
}

func TestExampleTwoThree(t *testing.T) {
	exp := []int{0, 1}
	res := twoSum([]int{3, 3}, 6)

	if !slices.Contains(res, exp[0]) || !slices.Contains(res, exp[1]) {
		t.Errorf(`twoSum result did not contain %d and %d!`, exp[0], exp[1])
	}
}

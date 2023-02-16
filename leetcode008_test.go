package leetcode

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Search in rotated array test 1", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		expectedResult := 4
		output := Search(nums, target)
		if expectedResult != output {
			t.Errorf("Error\nExpected %d\nOutput %d", expectedResult, output)
		}
	})
	t.Run("Search in rotated array test 1", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 3
		expectedResult := -1
		output := Search(nums, target)
		if expectedResult != output {
			t.Errorf("Error\nExpected %d\nOutput %d", expectedResult, output)
		}
	})
	t.Run("Search in rotated array test 1", func(t *testing.T) {
		nums := []int{1}
		target := 0
		expectedResult := -1
		output := Search(nums, target)
		if expectedResult != output {
			t.Errorf("Error\nExpected %d\nOutput %d", expectedResult, output)
		}
	})
}

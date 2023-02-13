package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	t.Run("Contains Duplicate test: 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		expectedResult := true
		results := ContainsDuplicate(nums)
		if !results == expectedResult {
			t.Errorf("Error\nExpected: %t\nFound: %t", expectedResult, results)
		}
	})
	t.Run("Contains Duplicate test: 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expectedResult := false
		results := ContainsDuplicate(nums)
		if !results == expectedResult {
			t.Errorf("Error\nExpected: %t\nFound: %t", expectedResult, results)
		}
	})
	t.Run("Contains Duplicate test: 3", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		expectedResult := true
		results := ContainsDuplicate(nums)
		if !results == expectedResult {
			t.Errorf("Error\nExpected: %t\nFound: %t", expectedResult, results)
		}
	})
}

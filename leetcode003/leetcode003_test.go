package leetcodego

import "testing"

func TestContainsDuplicate(t *testing.T) {

	t.Run("Contains Duplicate test: 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		results := ContainsDuplicate(nums)
		expectedResult := true
		if !results {
			t.Errorf("Error\nExpected: %t\nOutput: %t", expectedResult, results)
		}
	})
	t.Run("Contains Duplicate test: 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		results := ContainsDuplicate(nums)
		expectedResult := false
		if results {
			t.Errorf("Error\nExpected: %t\nOutput: %t", expectedResult, results)
		}
	})
	t.Run("Contains Duplicate test: 3", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		results := ContainsDuplicate(nums)
		expectedResult := true
		if !results {
			t.Errorf("Error\nExpected: %t\nOutput: %t", expectedResult, results)
		}
	})
}

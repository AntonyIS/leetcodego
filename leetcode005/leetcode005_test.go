package leetcodego

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Run("MaxSubArray Test 1", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		expectedResults := 6
		output := MaxSubArray(nums)
		if output != expectedResults {
			t.Errorf("\nExpected: %d\nOutput: %d", expectedResults, output)
		}
	})
	t.Run("MaxSubArray Test 2", func(t *testing.T) {
		nums := []int{1}
		expectedResults := 1
		output := MaxSubArray(nums)
		if output != expectedResults {
			t.Errorf("\nExpected: %d\nOutput: %d", expectedResults, output)
		}
	})
	t.Run("MaxSubArray Test 3", func(t *testing.T) {
		nums := []int{5, 4, -1, 7, 8}
		expectedResults := 23
		output := MaxSubArray(nums)
		if output != expectedResults {
			t.Errorf("\nExpected: %d\nOutput: %d", expectedResults, output)
		}
	})
}

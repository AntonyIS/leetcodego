package leetcode

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	t.Run("Max Product test 1", func(t *testing.T) {
		nums := []int{2, 3, -2, 4}
		expectedResult := 6
		output := MaxProduct(nums)
		if expectedResult != output {
			t.Errorf("Error\nExpected: %d\nOutput: %d", expectedResult, output)
		}
	})
	t.Run("Max Product test 2", func(t *testing.T) {
		nums := []int{-2,0,-1}
		expectedResult := 0
		output := MaxProduct(nums)
		if expectedResult != output {
			t.Errorf("Error\nExpected: %d\nOutput: %d", expectedResult, output)
		}
	})
}

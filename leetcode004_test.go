package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	t.Run("Product Except Self 10", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expectedResults := []int{24, 12, 8, 6}
		results := ProductExceptSelf(nums)
		if !reflect.DeepEqual(results, expectedResults) {
			t.Errorf("\nExpected: %d\nOutput: %d", expectedResults, results)
		}
	})
	t.Run("Product Except Self 10", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		expectedResults := []int{0, 0, 9, 0, 0}
		results := ProductExceptSelf(nums)
		if !reflect.DeepEqual(results, expectedResults) {
			t.Errorf("\nExpected: %d\nOutput: %d", expectedResults, results)
		}
	})
}

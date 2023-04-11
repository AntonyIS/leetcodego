package leetcodego

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("TwoSum Test: 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		expectedResult := []int{0, 1}
		target := 9
		results := TwoSum(nums, target)
		if !reflect.DeepEqual(results, expectedResult) {
			t.Errorf("Error: %d + %d != %d", results[0], results[1], target)
		}
	})

	t.Run("TwoSum Test: 2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		expectedResult := []int{1, 2}
		target := 6
		results := TwoSum(nums, target)
		if !reflect.DeepEqual(results, expectedResult) {
			t.Errorf("Error: %d + %d != %d", results[0], results[1], target)
		}
	})

}

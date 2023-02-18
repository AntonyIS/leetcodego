package leetcode

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	t.Run("Three sum test 1", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		expectedResult := [][]int{{-1, -1, 2}, {-1, 0, 1}}
		output := ThreeSum(nums)
		if !reflect.DeepEqual(expectedResult, output) {
			t.Errorf("Error\nExpected %d\nOutput %d", expectedResult, output)
		}
	})
	t.Run("Three sum test 2", func(t *testing.T) {
		nums := []int{0, 1, 1}
		expectedResult := [][]int{}
		output := ThreeSum(nums)
		if !reflect.DeepEqual(expectedResult, output) {
			t.Errorf("Error\nExpected %d\nOutput %d", expectedResult, output)
		}
	})
	t.Run("Three sum test 3", func(t *testing.T) {
		nums := []int{0,0,0}
		expectedResult := [][]int{{0,0,0}}
		output := ThreeSum(nums)
		if !reflect.DeepEqual(expectedResult, output) {
			t.Errorf("Error\nExpected %d\nOutput %d", expectedResult, output)
		}
	})

}

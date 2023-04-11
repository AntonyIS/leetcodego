package leetcodego

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	t.Run("Three sum test 1", func(t *testing.T) {
		nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		expectedResult := 49
		output := MaxArea(nums)
		if expectedResult != output {
			t.Errorf("Error\nExpected %d\nOutput %d\n", expectedResult, output)
		}
	})
	t.Run("Three sum test 2", func(t *testing.T) {
		nums := []int{1, 1}
		expectedResult := 1
		output := MaxArea(nums)
		if expectedResult != output {
			t.Errorf("Error\nExpected %d\nOutput %d\n", expectedResult, output)
		}
	})

}

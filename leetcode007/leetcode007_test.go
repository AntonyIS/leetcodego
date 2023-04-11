package leetcodego
import (
	"testing"
)

func TestFindMin(t *testing.T) {
	t.Run("Find Min test 1", func(t *testing.T) {
		nums := []int{3, 4, 5, 1, 2}
		expectedResult := 1
		output := FindMin(nums)
		if output != expectedResult {
			t.Errorf("Error\nExpected %d\nOuput %d", expectedResult, output)
		}
	})
	t.Run("Find Min test 2", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		expectedResult := 0
		output := FindMin(nums)
		if output != expectedResult {
			t.Errorf("Error\nExpected %d\nOuput %d", expectedResult, output)
		}
	})
	t.Run("Find Min test 3", func(t *testing.T) {
		nums := []int{11, 13, 15, 17}
		expectedResult := 11
		output := FindMin(nums)
		if output != expectedResult {
			t.Errorf("Error\nExpected %d\nOuput %d", expectedResult, output)
		}
	})
}

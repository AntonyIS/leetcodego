package leetcode011

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Run("Climbing stairs test 1", func(t *testing.T) {
		n := 2
		expectedResult := 2
		output := ClimbStairs(n)
		if output != expectedResult {
			t.Errorf("Error\nExpected %d\nOutput %d\n", expectedResult, output)
		}

	})

	t.Run("Climbing stairs test 2", func(t *testing.T) {
		n := 3
		expectedResult := 3
		output := ClimbStairs(n)
		if output != expectedResult {
			t.Errorf("Error\nExpected %d\nOutput %d\n", expectedResult, output)
		}
	})
}

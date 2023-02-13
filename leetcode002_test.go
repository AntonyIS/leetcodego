package leetcode

import "testing"

func TestBestTimeToBuyStock(t *testing.T) {
	t.Run("Best time to buy stock test: 1", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		expectedProfit := 5
		results := BestTimeToBuyStock(prices)

		if results != expectedProfit {
			t.Errorf("\nExpected: %d\nOutput: %d", results, expectedProfit)
		}
	})
	t.Run("Best time to buy stock test: 2", func(t *testing.T) {
		prices := []int{7, 6, 4, 3, 1}
		expectedProfit := 0
		results := BestTimeToBuyStock(prices)

		if results != expectedProfit {
			t.Errorf("\nExpected: %d\nOutput: %d", results, expectedProfit)
		}
	})
}

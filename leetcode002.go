package leetcode

/*
121. Best Time to Buy and Sell Stock
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/

func BestTimeToBuyStock(prices []int) int {
	// Assign left,right and maxProfit pointers of the prices array
	left, right, maxProfit := 0, 1, 0
	// Iterate over pricess as long as right is less that len prices
	for right < len(prices) {
		// Check if right value is greater than left to get max value
		if prices[left] < prices[right] {
			// Find the difference(profit) between the two pointers
			profit := prices[right] - prices[left]
			// Compare current profit with maxProfit
			if profit > maxProfit {
				// Change maxProfit value to profit
				maxProfit = profit
			}
		} else {
			// reassign the left pointer
			left = right
		}
		// Move the right pointer forward
		right += 1
	}
	return maxProfit
}


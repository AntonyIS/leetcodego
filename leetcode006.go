package leetcode

/*
152. Maximum Product Subarray
Given an integer array nums, find a  subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.
Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

func MaxProduct(nums []int) int {
	maxProduct, curMax, curMin := nums[0], 1, 1

	for _, num := range nums {
		// The 0 edge case, reset curMax and curMin to i
		if num == 0 {
			curMax, curMin = 1, 1
		}
		tmp := curMax * num
		curMax = func(args ...int) int {
			res := args[0]
			for _, i := range args {
				if i > res {
					res = i
				}
			}
			return res
		}(curMax*num, curMin*num, num)

		curMin = func(args ...int) int {
			res := args[0]
			for _, i := range args {
				if i < res {
					res = i
				}
			}
			return res
		}(tmp, curMin*num, num)

		maxProduct = func(a, b int) int {
			if a > b {
				return a
			}
			return b

		}(maxProduct, curMax)

	}
	return maxProduct
}

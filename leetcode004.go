package leetcode

/*
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

func ProductExceptSelf(nums []int) []int {
	// Create array to return the results
	res := make([]int, len(nums))
	// Assign pointers to start and end of nums
	prefix, postfix := 1, 1
	// Iterate over nums
	for i := 0; i < len(nums); i++ {
		// Assign prefix at index i
		res[i] = prefix
		// Find products from left to right
		prefix *= nums[i]
	}
	// Iterate over nums
	for j := len(nums) - 1; j >= 0; j-- {
		// Assign postfix at index j
		res[j] *= postfix
		// Find products from right to left
		postfix *= nums[j]
	}
	return res
}

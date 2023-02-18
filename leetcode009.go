package leetcode

import "sort"

/*
15. 3Sum
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

func ThreeSum(nums []int) [][]int {
	// define slice to store results
	res := [][]int{}
	// in place sort
	sort.Ints(nums)
	// loop through values
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		// define right and left pointers
		l, r := i+1, len(nums)-1
		// Get the current sum
		for l < r {
			sum := v + nums[l] + nums[r]
			// check if s > or < 0
			if sum > 0 {
				// move right pointer to the left
				r -= 1
			} else if sum < 0 {
				// move left pointer to the right
				l += 1
			} else {
				// add values in res slice
				res = append(res, []int{v, nums[l], nums[r]})
				// move left pointer to the right
				l += 1
				// check if left pointer valus is being repeated
				for nums[l] == nums[l-1] && l < r {
					// move left pointer to the right
					l += 1
				}
			}
		}
	}
	return res
}

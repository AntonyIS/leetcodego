package leetcodego

/*
153. Find Minimum in Rotated Sorted Array
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.



Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Example 3:
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.
*/

func FindMin(nums []int) int {
	// Min value
	res := nums[0]
	// Left and right pointer
	left, right := 0, len(nums)-1
	// Loop through values in nums
	for left <= right {
		// Check if nums is sorted
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			break
		} else {
			// Get the mid point
			mid := (left + right) / 2
			// Check if the mid value(nums[mid]) is less than the current min(res)
			res = min(res, nums[mid])
			// Check res is in the left or right section of nums
			if nums[mid] >= nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

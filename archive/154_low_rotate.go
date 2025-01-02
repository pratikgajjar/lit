package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func binary(nums []int, low, high int) int {
	if low == high {
		return nums[low]
	}
	mid := low + ((high - low) / 2)

	if nums[low] > nums[mid] {
		return binary(nums, low, mid)
	}
	if nums[mid] > nums[high] {
		return binary(nums, mid+1, high)
	}
	if nums[mid] == nums[high] {
		return binary(nums, low, high-1)
	}
	return nums[low]
}

func findMin(nums []int) int {
	switch len(nums) {
	case 0:
		return -1
	case 1:
		return nums[0]
	case 2:
		return min(nums[0], nums[1])
	}
	return binary(nums, 0, len(nums)-1)
}

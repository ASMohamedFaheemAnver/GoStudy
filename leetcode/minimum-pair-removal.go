package leetcode

import (
	"math"
)

func minimumPairRemoval(nums []int) int {
	switches := 0
	for len(nums) > 1 {
		ascending := true
		lefIndex := 0
		minSum := math.MaxInt
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				ascending = false
			}
			if nums[i]+nums[i+1] < minSum {
				lefIndex = i
				minSum = nums[i] + nums[i+1]
			}
		}
		if ascending {
			return switches
		}
		switches++
		nums[lefIndex] = minSum
		nums = append(nums[:lefIndex+1], nums[lefIndex+2:]...)
	}

	return switches
}

func MinimumPairRemoval(nums []int) int {
	return minimumPairRemoval(nums)
}

package leetcode

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if j, exists := seen[target-num]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}
	return []int{}
}

// Alternative: Brute force O(n²) solution
func twoSumBruteForce(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

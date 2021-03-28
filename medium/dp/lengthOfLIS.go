package dp

// Longest Increasing Subsequence

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]float64, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	// 关键是谁追谁
	for j := 0; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[j] > nums[i] {
				dp[j] = math.Max(dp[j], dp[i]+1)
			}
		}
	}
	var maxV float64
	for i := 0; i < len(nums); i++ {
		if dp[i] > maxV {
			maxV = dp[i]
		}
	}
	return int(maxV)
}

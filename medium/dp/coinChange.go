package dp

import "math"

func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	return dpCoinChange(memo, coins, amount)
}

func dpCoinChange(memo map[int]int, coins []int, amount int) int {
	// 读缓存
	if v, ok := memo[amount]; ok {
		return v
	}
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// 设置极值
	res := math.MaxFloat64

	// 选择一个硬币
	for _, c := range coins {
		subproblem := dpCoinChange(memo, coins, amount-c)
		if subproblem == -1 {
			continue
		}
		res = math.Min(float64(subproblem+1), res)
	}

	// 没有解
	if res == math.MaxFloat64 {
		memo[amount] = -1
		return -1
	}

	// 写缓存
	memo[amount] = int(res)
	return memo[amount]
}
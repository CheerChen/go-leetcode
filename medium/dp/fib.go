package dp

func fib(N int) int {
	memo := make(map[int]int)
	return dpFib(N, memo)
}

func dpFib(N int, memo map[int]int) int {
	// memo
	if memo[N] != 0 {
		return memo[N]
	}
	// basecase
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}

	// select
	memo[N] = dpFib(N-1, memo) + dpFib(N-2, memo)
	return memo[N]
}
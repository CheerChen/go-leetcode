package hard

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func trap(heightMap []int) int {
	if len(heightMap) == 0 {
		return 0
	}
	n := len(heightMap)
	lmax := make(map[int]int)
	rmax := make(map[int]int)

	lmax[0] = heightMap[0]
	rmax[n-1] = heightMap[n-1]
	// 从左向右递增
	for i := 1; i < n; i++ {
		lmax[i] = max(heightMap[i], lmax[i-1])
	}
	// 从右向左递增
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(heightMap[i], rmax[i+1])
	}
	var vol int
	for i,h := range heightMap {
		vol += min(lmax[i], rmax[i])-h
	}
	return vol
}
sort.Slice(a, func(i, j int) bool {
	return a[i] > a[j] // 降順にソート
})

func abs(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func reverse(s string) string {
	runes := []rune(s) // rune型はcode pointを単位として扱う型
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(min int, i int) int {
	if min > i {
		min = i
	}
	return min
}

func max(max int, i int) int {
	if max < i {
		max = i
	}
	return max
}

func tMax(t0 int, t1 int, t2 int) int {
	max := -1
	if t0 < t1 {
		if t1 < t2 {
			max = t2
		} else {
			max = t1
		}
	} else {
		if t0 < t2 {
			max = t2
		} else {
			max = t0
		}
	}
	return max
}

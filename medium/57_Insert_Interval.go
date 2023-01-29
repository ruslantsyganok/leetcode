func insert(intervals [][]int, newInterval []int) [][]int {

	var res [][]int

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if intervals[i][1] < newInterval[0]  {
			res = append(res, intervals[i])
		} else {
            newInterval = []int{min(intervals[i][0], newInterval[0]), max(intervals[i][1], newInterval[1])}
		}
	}

	res = append(res, newInterval)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

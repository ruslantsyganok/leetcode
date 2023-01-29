func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    var counter int 
    end := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        if  end <= intervals[i][0] {
            end = intervals[i][1]
        } else {
            counter++
            end = min(end, intervals[i][1])
        }
    }

    return counter 
}

func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}

func permute(nums []int) [][]int {
	return generatePermutations([]int{}, nums)
}

func generatePermutations(curr, remains []int) [][]int {
	if len(remains) == 0 {
		return [][]int{curr}
	}

	var res [][]int

	for i := 0; i < len(remains); i++ {
		c, r := make([]int, len(curr)), make([]int, len(remains))
		copy(c, curr)
		copy(r, remains)

		curr = append(curr, remains[i])
		res = append(res, generatePermutations(curr, append(append(remains[:i]), remains[i+1:]...))...)

		curr = c
		remains = r
	}

	return res
}

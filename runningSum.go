package running_sum

func runningSum(nums []int) []int {
	ret := make([]int, len(nums))
	accum := 0
	for idx, val := range nums {
		accum += val
		ret[idx] = accum
	}
	return ret
}

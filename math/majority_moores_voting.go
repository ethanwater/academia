package math


func MooresVoting(nums []int) int {
	count, candidate := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		switch {
		case num == candidate:
			count++
		default:
			count--
		}
	}

	return candidate
}

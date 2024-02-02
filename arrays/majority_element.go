package arrays

func MajorityElement(nums []int) int {
	count, canditate := 0, 0
	for _, num := range nums {
		if count == 0 {
			canditate = num
		}
		switch {
		case num == canditate:
			count++
		default:
			count--
		}
	}
	return canditate
}

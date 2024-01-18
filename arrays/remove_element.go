package arrays

func RemoveElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}

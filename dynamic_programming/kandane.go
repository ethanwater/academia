package dynamic_programming

func KandanesMaxDifference(nums []int) int {
	min := nums[0] //we set the default minimum to the first element as a holder, likely will change in iteration
	diff := 0 //default difference is 0

	for i:=1; i<len(nums); i++ {
		calc := nums[i] - min
		
		//if we detect a smaller number than the minimum, we make it the new minimum
		if min > nums[i] {
			min = nums[i]
			continue //no need to do rest of logic, so just continue here
		}

		if calc > diff {
			diff = calc
		}
	}

	return diff
}

func KandanesMaxSubarraySum(nums []int) int {
	sum, curr := nums[0], nums[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i:=0; i<len(nums); i++ {
		curr = max(nums[i], curr+nums[i])
		sum = max(sum, curr)	
	}

	return sum
}

package arrays

func MaxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			temp := prices[i] - prices[j]
			if temp > profit {
				profit = temp
			}
		}
	}

	return profit
}

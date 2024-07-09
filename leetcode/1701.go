func averageWaitingTime(customer [][]int) float64 {
    totalWait := customer[0][0] + customer[0][1]
    sum := customer[0][1]

	for i := 1; i < len(customer); i++ {
        if totalWait-customer[i][0] >= 0 {
            totalWait += customer[i][1]
            sum += totalWait-customer[i][0]
        } else {
            totalWait = customer[i][1] + customer[i][0]
            sum += customer[i][1]
        }
	}

    return float64(sum)/float64(len(customer))
}

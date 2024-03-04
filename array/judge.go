func findJudge(n int, trust [][]int) int {
	castvote := vote(trust)
	for _, villager := range trust {
		if villager[1] == 0 && castvote {
			return villager[1]
		}
	}
	return -1
}

func vote(trust [][]int) bool {
	judge := trust[0][1]
	
	for _, villager := range trust {
		if villager[1] != judge {
			return false
		}
		if villager[0] == judge {
			return false
		}
	}

	return true
}

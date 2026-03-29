package begginer

func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, account := range accounts {
		totalAccount := 0

		for _, money := range account {
			totalAccount += money
		}

		maxWealth = max(totalAccount, maxWealth)
	}

	return maxWealth
}

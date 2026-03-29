package begginer

func RunningSum(nums []int) []int {
	sums := []int{}
	for i, value := range nums {
		previousSums := 0
		if i != 0 {
			previousSums = sums[i-1]
		}

		sums = append(sums, value+previousSums)
	}

	return sums
}

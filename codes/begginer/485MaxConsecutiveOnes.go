package begginer

func FindMaxConsecutiveOnes(nums []int) int {
	var maxConsecutiveOnes, consecutiveOnes int
	for _, number := range nums {
		if number != 1 {
			maxConsecutiveOnes = max(maxConsecutiveOnes, consecutiveOnes)
			consecutiveOnes = 0
			continue
		}

		consecutiveOnes++
	}

	return max(maxConsecutiveOnes, consecutiveOnes)
}

package begginer

func FindNumbers(nums []int) int {
	var evenDigits int = 0
	for _, number := range nums {
		digits := 0
		for {
			digits++
			number = number / 10
			if number == 0 {
				break
			}
		}

		if digits%2 == 0 {
			evenDigits++
		}
	}
	
	return evenDigits
}

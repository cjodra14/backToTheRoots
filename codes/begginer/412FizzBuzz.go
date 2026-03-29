package begginer

import "fmt"

func FizzBuzz(n int) []string {
	index := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			index[i-1] = "FizzBuzz"

			continue
		}

		if i%3 == 0 {
			index[i-1] = "Fizz"

			continue
		}

		if i%5 == 0 {
			index[i-1] = "Buzz"

			continue
		}

		index[i-1] = fmt.Sprintf("%v", i)
	}

	return index
}

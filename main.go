package main

import (
	"fmt"

	"github.com/cjodra14/backToTheRoots/codes/begginer"
)

func main() {
	fmt.Println(begginer.RunningSum([]int{1, 2, 3, 4}))
	fmt.Println(begginer.MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(begginer.FizzBuzz(15))
	fmt.Println(begginer.CanConstruct("aa", "ab"))
}

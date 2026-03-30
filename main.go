package main

import (
	"fmt"

	"github.com/cjodra14/backToTheRoots/codes/begginer"
)

func main() {
	fmt.Println(begginer.RunningSum([]int{1, 2, 3, 4}))
	fmt.Println(begginer.MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(begginer.FizzBuzz(15))
	fmt.Println(begginer.NumberOfSteps(14))
	n5 := &begginer.ListNode{Val: 5, Next: nil}
	n4 := &begginer.ListNode{Val: 4, Next: n5}
	n3 := &begginer.ListNode{Val: 3, Next: n4}
	n2 := &begginer.ListNode{Val: 2, Next: n3}
	n1 := &begginer.ListNode{Val: 1, Next: n2}

	fmt.Println(begginer.MiddleNode(n1))
	fmt.Println(begginer.CanConstruct("aa", "ab"))

	fmt.Println(begginer.FindMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
	fmt.Println(begginer.FindNumbers([]int{12, 345, 2, 6, 7896}))
}

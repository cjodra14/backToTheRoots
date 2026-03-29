package begginer

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	nodesList := make([]*ListNode, 0)

	for {
		nodesList = append(nodesList, head)
		if head.Next == nil {
			return nodesList[len(nodesList)/2]
		}
		head = head.Next
	}

	// FYI: this is ranked as the best solution, but I do not understand how it works
	// fast := head

	// for fast != nil {
	// 	fast = fast.Next
	// 	if fast != nil {
	// 		fast = fast.Next
	// 		head = head.Next
	// 	}
	// }

	// return head
}

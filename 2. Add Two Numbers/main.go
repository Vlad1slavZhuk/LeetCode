/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode = &ListNode{Val: 0, Next: nil}
	curr := ans
	addN := 0

	for l1 != nil || l2 != nil {
		var x, y int

		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		sum := addN + x + y
		addN = sum / 10
		curr.Next = &ListNode{Val: sum % 10, Next: nil}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if addN > 0 {
		curr.Next = &ListNode{Val: addN, Next: nil}
	}
	return ans.Next
}
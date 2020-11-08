# A Very Big Sum

Link to **LeetCode**: [LeetCode](https://leetcode.com/explore/featured/card/november-leetcoding-challenge/564/week-1-november-1st-november-7th/3516/)


> You can solve the problem in O(1) memory using bits operation. use shift left operation ( `<<` ) and or operation ( `|` ) to get the decimal value in one operation.

```go
func getDecimalValue(head *ListNode) int {
	result := 0

	for head != nil {
		result = result<<1 | head.Val
		head = head.Next
	}

	return result
}
```
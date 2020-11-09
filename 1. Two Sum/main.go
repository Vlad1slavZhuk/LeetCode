// slow solution
func twoSum(nums []int, target int) []int {
	var ans []int

	for i := 0; i < len(nums)-1; i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}

	return ans
}

// fast solution
func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := nmap[target-v]; ok {
			return []int{j, i}
		}
		nmap[v] = i
	}
	return []int{-1, -1}
}

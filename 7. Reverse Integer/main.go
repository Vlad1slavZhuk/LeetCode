import "math"

func reverse(x int) int {
	var ans int

	for x != 0 {
		pop := x % 10
		x /= 10
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && pop < -8) {
			return 0
		}
		ans = ans*10 + pop
	}
	return ans
}

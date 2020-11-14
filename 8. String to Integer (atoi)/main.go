import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	var index int
	for i, ch := range s {
		if unicode.IsSpace(ch) {
			continue
		}
		if unicode.IsDigit(ch) || ch == '-' || ch == '+' {
			index = i
			break
		}
		return 0
	}

	neg := s[index] == '-'

	if s[index] == '-' || s[index] == '+' {
		index += 1
	}

	var sum int
	for _, ch := range s[index:] {
		if !unicode.IsNumber(ch) {
			break
		}

		sum *= 10
		sum += int(ch - '0')
		if neg && sum > math.MaxInt32+1 {
			return math.MinInt32
		} else if !neg && sum > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if neg {
		return -sum
	}
	return sum
}

package solutions

func DetermineFloors(str string) int {
	result := 0
	negativeOne := -1
	for i, char := range str {
		if char == '(' {
			char = '1'
		}
		if char == ')' {
			result = result + negativeOne
			continue
		}

		result = result + int(char-'0')

		if result == -1 {
			return i + 1
		}

	}
	return result
}

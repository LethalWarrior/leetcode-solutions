package plusone

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] <= 9 {
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

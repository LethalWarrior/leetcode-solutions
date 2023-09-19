package easy

func TwoSum(nums []int, target int) []int {
	diffs := map[int]int{}

	for i, num := range nums {
		diff := target - num
		if v, ok := diffs[diff]; ok {
			return []int{v, i}
		}
		diffs[num] = i
	}

	return []int{}
}

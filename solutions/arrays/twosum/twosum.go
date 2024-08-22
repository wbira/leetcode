package twosum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	var result []int
	for i, num := range nums {
		diff := target - num
		index, ok := m[diff]
		if ok && i != index {
			result = append(result, i, index)
			return result
		}
	}
	return nil
}

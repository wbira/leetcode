package twosum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		index, ok := m[diff]
		if ok && i != index {
			return []int{i, index}
		}
		m[num] = i
	}
	return nil
}

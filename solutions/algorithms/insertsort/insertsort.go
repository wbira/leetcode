package insertsort

func InstertSort(input []int) []int {

	for i := 1; i < len(input); i++ {
		j := i
		temp := input[j]

		for j > 0 && input[j-1] > temp {
			input[j] = input[j-1]
			j--
		}
		input[j] = temp
	}

	return input
}

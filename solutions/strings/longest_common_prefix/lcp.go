package longestcommonprefix

func addToSolution(partial string, char byte) string {
	byteSlice := []byte(partial)
	byteSlice = append(byteSlice, char)
	return string(byteSlice)
}

func LongestCommonPrefix(strs []string) string {
	var result = ""
	for i := 0; i <= 200; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return result
			}
		}
		result = addToSolution(result, strs[0][i])
	}
	return result
}

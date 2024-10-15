package searchstr

import "fmt"

func searchStr(haystack string, needle string) int {

	needleLen := len(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			j++
		} else {
			i = i - j
			j = 0
		}
		fmt.Printf("i %v j %v, len %v  i-j %v\n", i, j, needleLen, i-j)
		if needleLen == j {

			return i + 1 - j
		}
	}
	return -1
}

package roman

func resolveNextLetter(s string, i int) byte {
	if i < len(s) {
		return s[i]
	}
	return 0
}

func RomanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	for i := 0; i < len(s); i++ {
		currentLetter := s[i]
		nextLetter := resolveNextLetter(s, i+1)
		partialNum := romanValues[currentLetter]

		if romanValues[nextLetter] > partialNum {
			total -= partialNum
		} else {
			total += partialNum
		}

	}

	return total
}

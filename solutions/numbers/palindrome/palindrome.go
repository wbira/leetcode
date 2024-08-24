package palindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	curr := x
	reversed := 0
	for curr != 0 {
		digit := curr % 10
		reversed = 10*reversed + digit
		curr /= 10
	}
	return (x == reversed)
}

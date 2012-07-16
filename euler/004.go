package euler

import "strconv"

const Problem004Default = 1000

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(toCheck int) bool {
	stringVersion := strconv.Itoa(toCheck)
	return stringVersion == reverse(stringVersion)
}

func Problem004(limit int) int {
	largest := 0
	a := limit - 1
	b := limit - 1
	for a > 1 {
		if b < a {
			a--
			b = limit - 1
		} else {
			product := a * b
			if product > largest && isPalindrome(product) {
				largest = product
			}
			b--
		}
	}
	return largest
}

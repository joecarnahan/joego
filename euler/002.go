package euler

const Problem002Default = 4000000

func Problem002(limit int) int {
	total := 0
	first := 1
	second := 1
	next := first + second
	for next < limit {
		// Every third Fibonacci number is even, so we know that "next" is
		total += next
		// Move three positions, not just one
		for i := 0; i < 3; i++ {
			first = second
			second = next
			next = first + second
		}
	}
	return total
}

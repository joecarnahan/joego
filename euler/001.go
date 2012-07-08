package euler

const Problem001Default = 1000

func Problem001(limit int) int {
	multiples := []int{3, 5}
	product := multiples[0] * multiples[1]
	sum := 0

	// Sum up all multiples of 3 and multiples of 5
	for _, multiple := range multiples {
		current := multiple
		for current < limit {
			sum += current
			current += multiple
		}
	}
	// Then subtract all multiples of 15, as they were double-counted
	current := product
	for current < limit {
		sum -= current
		current += product
	}

	return sum
}

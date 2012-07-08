package euler

const Problem001Default = 1000

func Problem001(limit int) int {
	multiples := []int{3, 5}
	sum := 0
	for i := 0; i < limit; i++ {
		for _, multiple := range multiples {
			if i % multiple == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}

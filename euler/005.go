package euler

import "fmt"

const Problem005Default = 10

func mergeSlices(first, second []int) []int {
	// TODO
}

func collectFactors(toFactor int) []int {
	// TODO
}

func Problem005(limit int) int {
	// TODO
	primes := PrimesUnder(limit + 1)
	fmt.Println(primes)
	product := 1
	for _, prime := range primes {
		product *= prime
	}
	return product
}

package euler

import "container/list"

// Returns the lowest index in the given slice greater than the given index
// whose value is false, or returns the length of the slice if all values with
// index greater than the given index are true.
func nextPrime(lastPrime int, composites []bool) int {
	for i := lastPrime + 1; i < len(composites); i++ {
		if !composites[i] {
			return i
		}
	}
	return len(composites)
}

// Counts all of the values in the given slice whose index is 2 or greater and
// whose values are false.
func countPrimes(composites []bool) int {
	result := 0
	for i := 2; i < len(composites); i++ {
		if !composites[i] {
			result++
		}
	}
	return result
}

// Given a slice of Boolean values, return the indices of all false values
// where the index is 2 or greater. In other words, given a slice where
// composite indices are true and prime indices are false, return all of the
// prime indices.
func extractPrimes(composites []bool) []int {
	numPrimes := countPrimes(composites)
	result := make([]int, numPrimes)
	i := 0
	for j := 2; j < len(composites); j++ {
		if !composites[j] {
			result[i] = j
			i++
		}
	}
	return result
}

// Finds all prime numbers that are less than the given value.
func PrimesUnder(limit int) []int {
	composites := make([]bool, limit)
	prime := 2
	p := prime * prime
	for p < limit {
		for p < limit {
			composites[p] = true
			p += prime
		}
		prime = nextPrime(prime, composites)
		p = prime * prime
	}
	return extractPrimes(composites)
}

// Finds all prime factors of the given value.
func PrimeFactorsOf(toFactor int) *list.List {
	result := list.New()
	primes := PrimesUnder((toFactor + 2) / 2)
	for _, prime := range primes {
		for toFactor%prime == 0 {
			result.PushBack(prime)
			toFactor = toFactor / prime
		}
	}
	return result
}

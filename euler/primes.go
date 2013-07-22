package euler

// Returns the lowest index in the given slice greater than the given index
// whose value is false, or returns the length of the slice if all values with
// index greater than the given index are true.
func nextPrime(lastPrime int64, composites []bool) int64 {
	for i := lastPrime + 1; i < int64(len(composites)); i++ {
		if !composites[i] {
			return i
		}
	}
	return int64(len(composites))
}

// Counts all of the values in the given slice whose index is 2 or greater and
// whose values are false.
func countPrimes(composites []bool) int64 {
	result := 0
	for i := 2; i < len(composites); i++ {
		if !composites[i] {
			result++
		}
	}
	return int64(result)
}

// Given a slice of Boolean values, return the indices of all false values
// where the index is 2 or greater. In other words, given a slice where
// composite indices are true and prime indices are false, return all of the
// prime indices.
func extractPrimes(composites []bool) []int64 {
	numPrimes := countPrimes(composites)
	result := make([]int64, numPrimes)
	i := int64(0)
	for j := int64(2); j < int64(len(composites)); j++ {
		if !composites[j] {
			result[i] = j
			i++
		}
	}
	return result
}

// Finds all prime numbers that are less than the given value.
func PrimesUnder(limit int64) []int64 {
	// Build a Boolean slice where true values correspond to composite
	// numbers and false values correspond to prime numbers. Then,
	// convert that slice into a slice containing just the prime numbers.
	composites := make([]bool, limit)
	prime := int64(2)
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

// Returns a map from all prime factors of the given value to the number of
// times those factors appear in the prime factorization of the given value.
func PrimeFactorsOf(toFactor int64) map[int64]int64 {
	result := make(map[int64]int64)
	primes := PrimesUnder((toFactor + 2) / 2)
	for _, prime := range primes {
		for toFactor%prime == 0 {
			result[prime]++
			toFactor = toFactor / prime
		}
	}
	// If result is empty, then toFactor is its own prime factor.
	if len(result) == 0 {
		result[toFactor]++
	}
	return result
}

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

// Interface for generating successive prime numbers
type PrimeGenerator struct {
	// List should only contain values of type int
	knownPrimes *list.List
}

func NewPrimeGenerator() *PrimeGenerator {
	result := PrimeGenerator{list.New()}
	result.knownPrimes.PushFront(int(2))
	return &result
}

func divisibleByAny(toTest int, values *list.List) bool {
	for e := values.Front(); e != nil; e = e.Next() {
		if toTest % e.Value.(int) == 0 {
			return true
		}
	}
	return false
}

// Generate the next prime number. The first call to Next() on a given
// PrimeGenerator will return 2, then 3, then 5 and so on. Note that this will
// get really slow for larger prime numbers, and so this is implemented mainly
// as a curiosity.
func (generator *PrimeGenerator) Next() int {
	curr := generator.knownPrimes.Back().Value.(int)
	next := curr + 1
	for divisibleByAny(next, generator.knownPrimes) {
		next++
	}
	generator.knownPrimes.PushBack(next)
	return curr
}

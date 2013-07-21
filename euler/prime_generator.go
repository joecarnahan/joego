package euler

import "container/list"

// Interface for generating successive prime numbers
type PrimeGenerator struct {
	// List should only contain values of type int64
	knownPrimes *list.List
}

func NewPrimeGenerator() *PrimeGenerator {
	result := PrimeGenerator{list.New()}
	result.knownPrimes.PushFront(int64(2))
	return &result
}

func divisibleByAny(toTest int64, values *list.List) bool {
	for e := values.Front(); e != nil; e = e.Next() {
		if toTest%e.Value.(int64) == 0 {
			return true
		}
	}
	return false
}

// Generate the next prime number. The first call to Next() on a given
// PrimeGenerator will return 2, then 3, then 5 and so on. Note that this will
// get really slow for larger prime numbers, and so this is implemented mainly
// as a curiosity.
func (generator *PrimeGenerator) Next() int64 {
	curr := generator.knownPrimes.Back().Value.(int64)
	next := curr + 1
	for divisibleByAny(next, generator.knownPrimes) {
		next++
	}
	generator.knownPrimes.PushBack(next)
	return curr
}

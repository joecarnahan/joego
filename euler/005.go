package euler

import "container/list"

const Problem005Default = 20

func Problem005(limit int) int {
	factors := list.New()
	for i := 2; i <= limit; i++ {
		newFactors := PrimeFactorsOf(i)
		factors = MergeLists(factors, newFactors)
	}
	product := 1
	for e := factors.Front(); e != nil; e = e.Next() {
		product *= e.Value.(int)
	}
	return product
}

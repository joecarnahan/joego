package euler

import "container/list"

const Problem005Default = 20

func Problem005(limit int64) int64 {
	factors := list.New()
	for i := int64(2); i <= limit; i++ {
		newFactors := PrimeFactorsOf(i)
		factors = MergeLists(factors, newFactors)
	}
	product := int64(1)
	for e := factors.Front(); e != nil; e = e.Next() {
		product *= e.Value.(int64)
	}
	return product
}

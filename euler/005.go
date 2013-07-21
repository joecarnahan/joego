package euler

const Problem005Default = 10

// Adds all key-value pairs from "b" to "a". If a key appears in both maps,
// then value of that key in "a" is incremented by the value of that key in
// "b".
func mergeMaps(a, b map[int64]int64) {
	for key, value := range b {
		a[key] += value
	}
}


func Problem005(limit int64) int64 {
	// Map from factors to the number of times they are a factor
	factors := make(map[int64]int64)
	for i := int64(2); i <= limit; i++ {
		newFactors := PrimeFactorsMap(i)
		mergeMaps(factors, newFactors)
	}
	product := int64(1)
	for factor, count := range factors {
		for i := int64(0); i < count; i++ {
			product *= factor
		}
	}
	return product
}

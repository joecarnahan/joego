package euler

import "math"

const Problem003Default = 600851475143

func Problem003(toFactor int64) int64 {
	limit := int64(math.Sqrt(float64(toFactor)) + 0.5)
	primes := PrimesUnder(limit)
	for i := len(primes) - 1; i > 0; i-- {
		if toFactor%int64(primes[i]) == 0 {
			return int64(primes[i])
		}
	}
	return toFactor
}

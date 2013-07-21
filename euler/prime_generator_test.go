package euler

import (
	"testing"
)

func TestPrimeGenerator(t *testing.T) {
	generator := NewPrimeGenerator()
	for _, expected := range primes {
		if actual := generator.Next(); actual != expected {
			t.Errorf("Expected prime number %v, got %v", expected, actual)
		}
	}
}

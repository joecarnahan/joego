package euler

import (
	"reflect"
	"testing"
)

// The first 100 prime numbers.
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}

func TestPrimesUnder(t *testing.T) {
	type InputOutput struct {
		Input  int
		Output []int
	}
	testCases := []InputOutput{
		InputOutput{1, []int{}},
		InputOutput{2, []int{}},
		InputOutput{3, []int{2}},
		InputOutput{4, []int{2, 3}},
		InputOutput{10, []int{2, 3, 5, 7}},
		InputOutput{32, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}},
		InputOutput{primes[len(primes)-1] + 1, primes}}
	for _, testCase := range testCases {
		actual := PrimesUnder(testCase.Input)
		if !reflect.DeepEqual(actual, testCase.Output) {
			t.Errorf("For %v, got this:\n%v\nbut expected this:\n%v", testCase.Input,
				actual, testCase.Output)
		}
	}
}

func TestPrimesUnder(t *testing.T) {
	// TODO Once list equality is implemented and tested
}

func TestPrimeGenerator(t *testing.T) {
	generator := NewPrimeGenerator()
	for _, expected := range primes {
		if actual := generator.Next(); actual != expected {
			t.Errorf("Expected prime number %v, got %v", expected, actual)
		}
	}
}

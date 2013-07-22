package euler

import (
	"reflect"
	"testing"
)

// The first 100 prime numbers.
var primes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}

func TestPrimesUnder(t *testing.T) {
	type InputOutput struct {
		Input  int64
		Output []int64
	}
	testCases := []InputOutput{
		InputOutput{1, []int64{}},
		InputOutput{2, []int64{}},
		InputOutput{3, []int64{2}},
		InputOutput{4, []int64{2, 3}},
		InputOutput{10, []int64{2, 3, 5, 7}},
		InputOutput{32, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}},
		InputOutput{primes[len(primes)-1] + 1, primes}}
	for _, testCase := range testCases {
		actual := PrimesUnder(testCase.Input)
		if !reflect.DeepEqual(actual, testCase.Output) {
			t.Errorf("For %v, got this:\n%v\nbut expected this:\n%v", testCase.Input,
				actual, testCase.Output)
		}
	}
}

func mapSubset(a, b map[int64]int64) bool {
	for key, valueA := range a {
		valueB, present := b[key]
		if !present || (valueA != valueB) {
			return false
		}
	}
	return true
}

func mapsEqual(a, b map[int64]int64) bool {
	return mapSubset(a, b) && mapSubset(b, a)
}

func TestPrimeFactorsOf(t *testing.T) {
	type TestCase struct {
		Input  int64
		Output map[int64]int64
	}
	factorsOf3 := make(map[int64]int64)
	factorsOf3[int64(3)] = int64(1)
	factorsOf4 := make(map[int64]int64)
	factorsOf4[int64(2)] = int64(2)
	factorsOf18 := make(map[int64]int64)
	factorsOf18[int64(2)] = int64(1)
	factorsOf18[int64(3)] = int64(2)
	factorsOf1050 := make(map[int64]int64)
	factorsOf1050[int64(2)] = int64(1)
	factorsOf1050[int64(3)] = int64(1)
	factorsOf1050[int64(5)] = int64(2)
	factorsOf1050[int64(7)] = int64(1)
	testCases := []TestCase{
		TestCase{int64(3), factorsOf3},
		TestCase{int64(4), factorsOf4},
		TestCase{int64(18), factorsOf18},
		TestCase{int64(1050), factorsOf1050}}
	for _, testCase := range testCases {
		if factors := PrimeFactorsOf(testCase.Input); !mapsEqual(factors, testCase.Output) {
			t.Errorf("Expected factors of %v to be %v, got %v", testCase.Input, testCase.Output, factors)
		}
	}
}

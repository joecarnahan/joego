package euler

import (
	"testing"
	"reflect"
)

func TestPrimes(t *testing.T) {
	type InputOutput struct {
		Input int
		Output []int
	}
	testCases := []InputOutput{
		InputOutput{1, []int{}},
		InputOutput{2, []int{}},
		InputOutput{3, []int{2}},
		InputOutput{4, []int{2, 3}},
		InputOutput{10, []int{2, 3, 5, 7}},
		InputOutput{32, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}}}
	for _, testCase := range testCases {
		actual := PrimesUnder(testCase.Input)
		if !reflect.DeepEqual(actual, testCase.Output) {
			t.Errorf("For %v, got %v, expected %v", testCase.Input,
				actual, testCase.Output)
		}
	}
}

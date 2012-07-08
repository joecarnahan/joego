package euler

import "testing"

func TestProblem001(t *testing.T) {
	const expected = 42
	if out := Problem001(Problem001Default); out != expected {
		t.Errorf("Got %v, want %v", out, expected)
	}
}

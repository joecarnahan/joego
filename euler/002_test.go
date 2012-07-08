package euler

import "testing"

func TestProblem002(t *testing.T) {
	const expected = 4613732
	if out := Problem002(Problem002Default); out != expected {
		t.Errorf("Got %v, want %v", out, expected)
	}
}

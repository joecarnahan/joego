package euler

import "testing"

func TestProblem005(t *testing.T) {
	const expected = 232792560
	if out := Problem005(Problem005Default); out != expected {
		t.Errorf("Got %v, want %v", out, expected)
	}
}

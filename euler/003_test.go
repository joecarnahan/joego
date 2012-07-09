package euler

import "testing"

func TestProblem003(t *testing.T) {
	const expected int64 = 6857
	if out := Problem003(Problem003Default); out != expected {
		t.Errorf("Got %v, want %v", out, expected)
	}
}

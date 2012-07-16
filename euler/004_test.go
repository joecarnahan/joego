package euler

import "testing"

func TestProblem004(t *testing.T) {
	const expected int = 906609
	if out := Problem004(Problem004Default); out != expected {
		t.Errorf("Got %v, want %v", out, expected)
	}
}

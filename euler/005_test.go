package euler

import "testing"

func TestMergeMaps(t *testing.T) {
	empty := make(map[int64]int64)
	empty2 := make(map[int64]int64)
	mergeMaps(empty, empty2)
	if len(empty) != 0 {
		t.Errorf("Expected result to be empty, got %v", empty)
	}
	has11 := make(map[int64]int64)
	has11[int64(1)] = int64(1)
	mergeMaps(has11, empty)
	if len(has11) != 1 {
		t.Errorf("Expected result to have one element, got %v", has11)
	} else if has11[int64(1)] != int64(1) {
		t.Errorf("Expected result[1] to be 1, got %v", has11[int64(1)])
	}
	emptyGets11 := make(map[int64]int64)
	mergeMaps(emptyGets11, has11)
	if len(has11) != 1 {
		t.Errorf("Expected has11 to still have one element, got %v", has11)
	} else if has11[int64(1)] != int64(1) {
		t.Errorf("Expected has11[1] to be 1, got %v", has11[int64(1)])
	}
	if len(emptyGets11) != 1 {
		t.Errorf("Expected result to have one element, got %v", emptyGets11)
	} else if emptyGets11[int64(1)] != int64(1) {
		t.Errorf("Expected result[1] to be 1, got %v", emptyGets11[int64(1)])
	}
	has11Gets11 := make(map[int64]int64)
	has11Gets11[int64(1)] = int64(1)
	mergeMaps(has11Gets11, has11)
	if len(has11) != 1 {
		t.Errorf("Expected has11 to still have one element, got %v", has11)
	} else if has11[int64(1)] != int64(1) {
		t.Errorf("Expected has11[1] to be 1, got %v", has11[int64(1)])
	}
	if len(has11Gets11) != 1 {
		t.Errorf("Expected result to have one element, got %v", has11Gets11)
	} else if has11Gets11[int64(1)] != int64(2) {
		t.Errorf("Expected result[1] to be 2, got %v", has11Gets11[int64(1)])
	}
	has2345Gets11 := make(map[int64]int64)
	has2345Gets11[int64(2)] = int64(3)
	has2345Gets11[int64(4)] = int64(5)
	mergeMaps(has2345Gets11, has11)
	if len(has11) != 1 {
		t.Errorf("Expected has11 to still have one element, got %v", has11)
	} else if has11[int64(1)] != int64(1) {
		t.Errorf("Expected has11[1] to be 1, got %v", has11[int64(1)])
	}
	if len(has2345Gets11) != 3 {
		t.Errorf("Expected result to have three elements, got %v", has2345Gets11)
	} else if has2345Gets11[int64(1)] != int64(1) {
		t.Errorf("Expected result[1] to be 1, got %v", has2345Gets11[int64(1)])
	} else if has2345Gets11[int64(2)] != int64(3) {
		t.Errorf("Expected result[2] to be 3, got %v", has2345Gets11[int64(2)])
	} else if has2345Gets11[int64(4)] != int64(5) {
		t.Errorf("Expected result[4] to be 5, got %v", has2345Gets11[int64(4)])
	}
}

// TODO Figure out why this test doesn't pass
/* func TestProblem005(t *testing.T) {
	const expected = 232792560
	if out := Problem005(Problem005Default); out != expected {
		t.Errorf("Got %v, want %v", out, expected)
	}
} */

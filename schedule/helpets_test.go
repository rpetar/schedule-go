package schedule

import (
	"testing"
)

func TestTestEq(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	res := TestEq(a, b)

	if res != true {
		t.Error("Expected true, got ", res)
	}

	a = []int{1, 2, 3, 5}
	b = []int{1, 2, 3, 4}
	res = TestEq(a, b)

	if res != false {
		t.Error("Expected false, got ", res)
	}

	a = []int{1, 2, 3, 4, 5}
	b = []int{1, 2, 3, 4}
	res = TestEq(a, b)

	if res != false {
		t.Error("Expected false, got ", res)
	}
}

func TestCheckifSliceInSlices(t *testing.T) {
	tabu := [][]int{{1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 5}}
	val := []int{1, 2, 3, 4}
	res := CheckIfSliceInSlices(val, tabu)

	if res != true {
		t.Error("Expected true, got ", res)
	}

	val = []int{1, 2}
	res = CheckIfSliceInSlices(val, tabu)

	if res != false {
		t.Error("Expected false, got ", res)
	}

	val = []int{1, 2, 3, 4, 5}
	res = CheckIfSliceInSlices(val, tabu)

	if res != false {
		t.Error("Expected false, got ", res)
	}
}

func TestFilterSliceOfSlices(t *testing.T) {
	tabu := [][]int{{1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 5}}
	neighbors := [][]int{{1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 5}, {2, 3, 4}}
	res := FilterSliceOfSlices(neighbors, tabu)

	expected := [][]int{{2, 3, 4}}
	if len(res) != 1 {
		t.Error("Expected 1, got ", len(res))
	}
	for i, d := range res[0] {
		if d != expected[0][i] {
			t.Errorf("Expected slice element %d to be %d got %d", i, expected[0][i], d)
		}
	}
}

func TestUniqueSlices(t *testing.T) {
	neighbors := [][]int{{1, 2, 3}, {2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 5}}
	res := UniqueSlices(neighbors)

	if len(res) != 4 {
		t.Error("Expected 1, got ", len(res))
	}
}

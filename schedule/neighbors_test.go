package schedule

import (
	"math/rand"
	"testing"
)

func TestFindBestSolution(t *testing.T) {
	numDays := 28
	workingDays := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0}
	weights := []int{3, 4, 5, 4, 5, 6, 5, 6, 7, 6, 7, 8}
	blocks := [][]int{{2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}, {4, 1}, {4, 2}, {4, 3}, {5, 1}, {5, 2}, {5, 3}}
	minOff := 1
	maxOff := 3
	minWorking := 2
	maxWorking := 5
	eval := EvaluateFactory(numDays, workingDays, weights, blocks, minOff, maxOff, minWorking, maxWorking)

	// not real neighbors
	neighbors := [][]int{{10, 10, 10}, {9, 10, 10}, {11, 10, 10}, {10, 10}, {10, 9, 10}, {10, 11, 10}, {10, 10, 9}, {10, 10, 11}}

	res := FindBestSolution(neighbors, eval)
	if res.Score.Total != 9 {
		t.Error("Expected score to be 9, got ", res.Score.Total)
	}
}

func TestGetTotalWeight(t *testing.T) {
	solution := []int{0,1,2,3}
	weights := []int{10, 20, 10, 50}
	res := GetTotalWeight(weights, solution)
	if res != 90 {
		t.Error("Expected 90, got ", res)
	}
}

func TestGenRandomSolution(t *testing.T) {
	weights := []int{10, 20, 10, 50}
	maxWeight := 120
	r := rand.New(rand.NewSource(1))

	res := GenRandomSolution(maxWeight, weights, r)
	if res[0] != 1 && res[1] != 3 && res[2] != 3 {
		t.Error("Expected [1 3 3] result, got ", res)
	}
}

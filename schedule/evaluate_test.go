package schedule

import "testing"

func TestEval(t *testing.T) {
	solution := []int{10, 0, 1, 10, 10}
	numDays := 28
	workingDays := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0}
	weights := []int{3, 4, 5, 4, 5, 6, 5, 6, 7, 6, 7, 8}
	blocks := [][]int{{2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}, {4, 1}, {4, 2}, {4, 3}, {5, 1}, {5, 2}, {5, 3}}
	minOff := 1
	maxOff := 3
	minWorking := 2
	maxWorking := 5
	eval := EvaluateFactory(numDays, workingDays, weights, blocks, minOff, maxOff, minWorking, maxWorking)
	res := eval(solution)

	if res.Penalty != 0 {
		t.Error("Expected penalty to be 0, got ", res.Penalty)
	}

	if res.Total != -19 {
		t.Error("Expected total to be -19, got ", res.Total)
	}

	if res.Bonus != 19 {
		t.Error("Expected bonus to be 19, got ", res.Bonus)
	}

	solution = []int{11, 5, 10, 10}
	res = eval(solution)
	if res.Penalty != 0 {
		t.Error("Expected penalty to be 0, got ", res.Penalty)
	}

	if res.Total != -18 {
		t.Error("Expected total to be -19, got ", res.Total)
	}

	if res.Bonus != 18 {
		t.Error("Expected bonus to be 19, got ", res.Bonus)
	}

	solution = []int{5, 11, 11}
	res = eval(solution)
	if res.Penalty != 64 {
		t.Error("Expected penalty to be 64, got ", res.Penalty)
	}

	if res.Total != 51 {
		t.Error("Expected total to be 51, got ", res.Total)
	}

	if res.Bonus != 13 {
		t.Error("Expected bonus to be 13, got ", res.Bonus)
	}

}

func TestEvalWeight(t *testing.T) {
	solution := []int{10, 0, 1, 10, 10}
	numDays := 28
	weights := []int{3, 4, 5, 4, 5, 6, 5, 6, 7, 6, 7, 8}

	res := EvalWeight(numDays, weights, solution)

	if res != 0 {
		t.Error("Expected weight penlaty to be 0, got ", res)
	}

	solution = []int{5, 11, 11}

	res = EvalWeight(numDays, weights, solution)

	if res != 24 {
		t.Error("Expected weight penlaty to be 24, got ", res)
	}
}

func TestEvalDays(t *testing.T) {
	solution := []int{10, 0, 1, 10, 10}
	minOff := 1
	maxOff := 3
	minWorking := 2
	maxWorking := 5
	workingDays := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0}
	blocks := [][]int{{2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}, {4, 1}, {4, 2}, {4, 3}, {5, 1}, {5, 2}, {5, 3}}

	res := EvalDays(workingDays, blocks, solution, minOff, maxOff, minWorking, maxWorking)

	if res[0] != 0 {
		t.Error("Expected working days penlaty to be 0, got ", res[0])
	}

	if res[1] != 19 {
		t.Error("Expected working days bonus to be 0, got ", res[1])
	}
}

func TestGetDaysStats(t *testing.T) {
	solution := []int{10, 0, 1, 10, 10}

	workingDays := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0}
	blocks := [][]int{{2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}, {4, 1}, {4, 2}, {4, 3}, {5, 1}, {5, 2}, {5, 3}}

	consecutiveWorkings, consecutiveOffs, totalWorking, wrongDayOffPenalty := GetDaysStats(workingDays, blocks, solution)

	expectedConsecutiveWorkings := []int{5, 2, 2, 5, 5}
	expectedConsecutiveOffs := []int{2, 1, 2, 2, 2}

	if len(consecutiveWorkings) != len(expectedConsecutiveWorkings) {
		t.Error("Expected consecutive working length to be ", len(expectedConsecutiveWorkings), "got ", len(consecutiveWorkings))
	} else {
		for i, d := range consecutiveWorkings {
			if d != expectedConsecutiveWorkings[i] {
				t.Error("Expected consecutive working ", i, "to be ", expectedConsecutiveWorkings[i], "got ", d)
			}
		}
	}

	if len(consecutiveOffs) != len(expectedConsecutiveOffs) {
		t.Error("Expected consecutive offs length to be ", len(expectedConsecutiveOffs), "got ", len(consecutiveWorkings))
	} else {
		for i, d := range consecutiveOffs {
			if d != expectedConsecutiveOffs[i] {
				t.Error("Expected consecutive off ", i, "to be ", expectedConsecutiveOffs[i], "got ", d)
			}
		}
	}

	if totalWorking != 19 {
		t.Error("Expected totalWorking to be 19, got ", totalWorking)
	}
	if wrongDayOffPenalty != 0 {
		t.Error("Expected wrongDayOffPenalty to be 0, got ", wrongDayOffPenalty)
	}
}

func TestGetConsecutivePenalty(t *testing.T) {
	consecutive := []int{5, 2, 1, 6}
	res := GetConsecutivePenalty(consecutive, 2, 5)

	if res != 40 {
		t.Error("Expected penalty to be 40, got ", res)
	}

	consecutive = []int{5, 2, 3, 4}
	res = GetConsecutivePenalty(consecutive, 2, 5)

	if res != 0 {
		t.Error("Expected penalty to be 0, got ", res)
	}
}

func TestGetWrongDayOffPenalty(t *testing.T) {
	workingDays := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0}

	res := GetWrongDayOffPenalty(workingDays, 0, 5)
	if res != 0 {
		t.Error("Expected penalty to be 0, got ", res)
	}

	res = GetWrongDayOffPenalty(workingDays, 3, 5)
	if res != 80 {
		t.Error("Expected penalty to be 80, got ", res)
	}
}

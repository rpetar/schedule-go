package search

import (
	"github.com/Tymeshift/data-science-test/helpers"
	"github.com/Tymeshift/data-science-test/schedule"
	"math/rand"
	"testing"
)

// Random is expected to fail in most cases

func RandomTest(t *testing.T, params helpers.ProblemParams) {
	r := rand.New(rand.NewSource(1))

	evaluate, _, blocks := helpers.GetSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this

	initialSolution := schedule.GenRandomSolution(params.NumDays, blocks.Weights, r)
	solution := Random(evaluate, initialSolution, params, blocks, maxIterations, limitNotImproved, r)

	if solution.Score.Penalty != 0 {
		t.Skip("Expected solution penalty to be 0, got ", solution.Score.Penalty)
	}
}

func TestRandom(t *testing.T) {
	params := helpers.ParamSets[0]
	RandomTest(t, params)
}

func TestRandom1(t *testing.T) {
	params := helpers.ParamSets[1]
	RandomTest(t, params)
}

func TestRandom2(t *testing.T) {
	params := helpers.ParamSets[2]
	RandomTest(t, params)
}

func TestRandom3(t *testing.T) {
	params := helpers.ParamSets[3]
	RandomTest(t, params)
}

func BenchmarkRandom(b *testing.B) {
	r := rand.New(rand.NewSource(1))

	params := helpers.ParamSets[3]
	evaluate, _, blocks := helpers.GetSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this

	for i := 0; i < b.N; i++ {
		initialSolution := schedule.GenRandomSolution(params.NumDays, blocks.Weights, r)
		solution := Random(evaluate, initialSolution, params, blocks, maxIterations, limitNotImproved, r)
		if solution.Score.Penalty != 0 {
			b.Skip("Expected solution penalty to be 0, got ", solution.Score.Penalty)
		}
	}
}

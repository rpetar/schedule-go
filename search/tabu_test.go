package search

import (
	"github.com/Tymeshift/data-science-test/helpers"
	"github.com/Tymeshift/data-science-test/schedule"
	"math/rand"
	"testing"
)

func TabuSearchTest(t *testing.T, params helpers.ProblemParams) {
	r := rand.New(rand.NewSource(1))

	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	evaluate, findNeighborhood, blocks := helpers.GetSearchParams(params)
	tabuListSize := 1000
	initialSolution := schedule.GenRandomSolution(params.NumDays, blocks.Weights, r)
	res := TabuSearch(evaluate, findNeighborhood, initialSolution, maxIterations, tabuListSize, limitNotImproved)
	if res.Score.Penalty != 0 {
		t.Error("Expected solution penalty to be 0, got ", res.Score.Penalty)
	}
}

func TestTabuSearch(t *testing.T) {
	params := helpers.ParamSets[0]
	TabuSearchTest(t, params)
}

func TestTabuSearch2(t *testing.T) {
	params := helpers.ParamSets[1]
	TabuSearchTest(t, params)
}

func TestTabuSearch3(t *testing.T) {
	params := helpers.ParamSets[2]
	TabuSearchTest(t, params)
}

func TestTabuSearch4(t *testing.T) {
	params := helpers.ParamSets[3]
	TabuSearchTest(t, params)
}

func TestTabuSearch5(t *testing.T) {
	params := helpers.ParamSets[4]
	TabuSearchTest(t, params)
}

func BenchmarkTabuSearch(b *testing.B) {
	r := rand.New(rand.NewSource(1))

	params := helpers.ParamSets[2]
	evaluate, findNeighborhood, blocks := helpers.GetSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this

	for i := 0; i < b.N; i++ {
		tabuListSize := 1000
		initialSolution := schedule.GenRandomSolution(params.NumDays, blocks.Weights, r)
		solution := TabuSearch(evaluate, findNeighborhood, initialSolution, maxIterations, tabuListSize, limitNotImproved)
		if solution.Score.Penalty != 0 {
			b.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
		}
	}
}

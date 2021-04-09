package search

import (
	"github.com/Tymeshift/data-science-test/helpers"
	"github.com/Tymeshift/data-science-test/schedule"
	"math/rand"
)

// Random implements Random Walk
func Random(evaluate schedule.Evaluate, initialSolution []int, params helpers.ProblemParams, blocks schedule.BlocksData, maxIter, limitNotImproved int, r *rand.Rand) schedule.Solution {
	count := 0
	bestCost := evaluate(initialSolution)
	solution := schedule.Solution{
		Value: initialSolution,
		Score: bestCost,
	}
	bestSolution := solution

	notImprovedCounter := 0

	for count <= maxIter {

		nextSolutionValue := schedule.GenRandomSolution(params.NumDays, blocks.Weights, r)
		nextSolutionScore := evaluate(nextSolutionValue)
		solution := schedule.Solution{
			Value: nextSolutionValue,
			Score: nextSolutionScore,
		}

		cost := bestSolution.Score.Total - solution.Score.Total
		if cost >= 0 && solution.Score.Penalty <= bestSolution.Score.Penalty {
			notImprovedCounter = 0
			bestSolution = solution
		} else {
			notImprovedCounter++
		}

		if bestSolution.Score.Penalty == 0 && notImprovedCounter >= limitNotImproved {
			return bestSolution
		}
		count++
	}
	return bestSolution
}

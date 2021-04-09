package search

import (
	"github.com/Tymeshift/data-science-test/schedule"
)

// TabuSearch - performs the search using the tabu search technique
func TabuSearch(evaluate schedule.Evaluate, findNeighborhood schedule.FindNeighborhood, initialSolution []int, maxIter, tabuSize, limitNotImproved int) schedule.Solution {
	count := 0
	bestCost := evaluate(initialSolution)
	solution := schedule.Solution{
		Value: initialSolution,
		Score: bestCost,
	}
	bestSolution := solution
	tabuList := make([][]int, tabuSize)

	notImprovedCounter := 0

	for count <= maxIter {
		// get all of the neighbors
		neighbors := findNeighborhood(solution.Value)
		neighbors = schedule.FilterSliceOfSlices(neighbors, tabuList)

		if len(neighbors) > 0 {
			solution = schedule.FindBestSolution(neighbors, evaluate)
			// if the new solution is better,
			// update the current solution with the new solution
			if solution.Score.Penalty < bestSolution.Score.Penalty || (solution.Score.Penalty == bestSolution.Score.Penalty && solution.Score.Total < bestSolution.Score.Total) {
				notImprovedCounter = -1
				bestSolution = solution
			}
			notImprovedCounter++

			if bestSolution.Score.Penalty == 0 && notImprovedCounter >= limitNotImproved {
				return bestSolution
			}

			tabuVal := make([]int, len(solution.Value))
			copy(tabuVal, solution.Value)
			tabuList = append(tabuList, tabuVal)

			if len(tabuList) > tabuSize {
				tabuList = tabuList[1:]
			}

		}
		count++
	}

	return bestSolution
}

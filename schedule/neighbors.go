package schedule

import (
	"math/rand"
)

// FindBestSolution finds the best neighbor
func FindBestSolution(neighbors [][]int, evaluate Evaluate) Solution {
	evaluated := make([]Solution, len(neighbors))
	for i := range neighbors {
		evaluated[i] = Solution{
			Value: neighbors[i],
			Score: evaluate(neighbors[i]),
		}
	}

	min := evaluated[0]
	for _, e := range evaluated {
		if e.Score.Penalty < min.Score.Penalty || (e.Score.Penalty == min.Score.Penalty && e.Score.Total < min.Score.Total) {
			min = e
		}
	}
	return min
}

// GetTotalWeight calculates total weight of the solution
func GetTotalWeight(weights []int, solution []int) int {
	totalWeight := 0
	for _, s := range solution {
		totalWeight += weights[s]
	}
	return totalWeight
}

// GenRandomSolution generates random solution, to be used as the initial solution for the search
func GenRandomSolution(maxWeight int, weights []int, r *rand.Rand) []int {
	ir := IntRange{Min: 0, Max: len(weights) - 1}
	curWeight := 0
	var blocks []int
	for maxWeight > curWeight {
		randBlock := ir.NextRandom(r)
		curWeight += weights[randBlock]
		if curWeight <= maxWeight {
			blocks = append(blocks, randBlock)
		}
	}

	return blocks
}

// FindNeighborhoodFactory creates a FindNeighborhood function to be used in a search
// FindNeighborhood should return all possible unique neighbors of a given solution.
func FindNeighborhoodFactory(dataLength int) FindNeighborhood {
	return func(solution []int) [][]int {
		neighbors := make([][]int, 0)

		for i, s := range solution{
		//	Increase index of i-th block by 1
			newNeighborInc := make([]int, len(solution))
			copy(newNeighborInc, solution)
			if (s+1) < dataLength{
				newNeighborInc[i] = s + 1
				neighbors = append(neighbors, newNeighborInc)
			}

			// Decrease index of i-th block by 1
			newNeighborDec := make([]int, len(solution))
			copy(newNeighborDec, solution)
			if (s - 1) >= 0 {
				newNeighborDec[i] = s - 1
				neighbors = append(neighbors, newNeighborDec)
			}

			// Remove block
			newNeighborRemove := make([]int, len(solution))
			copy(newNeighborRemove, solution)
			newNeighborRemove = append(newNeighborRemove[:i], newNeighborRemove[i+1:]...)
			neighbors = append(neighbors, newNeighborRemove)
	}

	for i := 0; i < dataLength; i++ {
		newNeighborAdd := make([]int, len(solution))
		copy(newNeighborAdd, solution)
		newNeighborAdd = append(newNeighborAdd, i)
		neighbors = append(neighbors, newNeighborAdd)
	}

	return neighbors
	}
}

// BlocksData contains the blocks and associated parameters that are used to constract the solution. Inspired by knapsack problem
type BlocksData struct {
	Weights []int
	Values  []int
	Blocks  [][]int
}

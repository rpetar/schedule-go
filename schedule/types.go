package schedule

// Solution - type that containes solution data and score
type Solution struct {
	Value []int
	Score SolutionScore
}

// SolutionScore - contains score type for the solution
type SolutionScore struct {
	Total   int
	Penalty int
	Bonus   int
}

// IntRange - min/max ranges
type IntRange struct {
	Min, Max int
}

// Evaluate - function that evaluates optimisation problem
type Evaluate func([]int) SolutionScore

// FindNeighborhood - function that finds neighbors for optimisation problem
type FindNeighborhood func([]int) [][]int


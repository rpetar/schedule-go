package schedule

// EvaluateFactory creates an Evaluate function to be used in a search to evaluate solutions
func EvaluateFactory(maxWeight int, days []int, weights []int, blocks [][]int, minOff int, maxOff int, minWorking int, maxWorking int) Evaluate {
	return func(solution []int) SolutionScore {
		penalty := 0
		weightPenalty := EvalWeight(maxWeight, weights, solution)
		penalty += weightPenalty
		workingPenalties := EvalDays(days, blocks, solution, minOff, maxOff, minWorking, maxWorking)
		penalty += workingPenalties[0]

		return SolutionScore{
			Total:   penalty - workingPenalties[1],
			Penalty: penalty,
			Bonus:   workingPenalties[1],
		}
	}
}

// EvalWeight evaluates solution weight
func EvalWeight(maxWeight int, weights []int, solution []int) int {
	totalWeight := GetTotalWeight(weights, solution)
	if totalWeight > maxWeight {
		return (totalWeight - maxWeight) * 8 // penalty for going over the limit
	}
	if totalWeight < maxWeight {
		return (maxWeight - totalWeight) * 4 // penalty for underperforming
	}
	return 0
}

// GetDaysStats gets different stats about the solution, i.e. consecutive, total working and penalty
func GetDaysStats(days []int, blocks [][]int, solution []int) ([]int, []int, int, int) {
	consecutiveWorkings := make([]int, 0, len(days))
	consecutiveOffs := make([]int, 0, len(days))
	totalWorking := 0
	curDay := 0
	wrongDayOffPenalty := 0

	for _, s := range solution {
		currentBlock := blocks[s]
		consecutiveWorkings = append(consecutiveWorkings, currentBlock[0])
		totalWorking += currentBlock[0]
		wrongDayOffPenalty += GetWrongDayOffPenalty(days, curDay, currentBlock[0])

		curDay += currentBlock[0]

		consecutiveOffs = append(consecutiveOffs, currentBlock[1])

		curDay += currentBlock[1]
	}
	return consecutiveWorkings, consecutiveOffs, totalWorking, wrongDayOffPenalty
}

// GetWrongDayOffPenalty checks if working day assigned to the pre-defined day-off and returns penalty
func GetWrongDayOffPenalty(days []int, curDay int, workingDays int) int {
	penalty := 0
	for i := 0; i < workingDays; i++ {
		if i+curDay >= len(days) {
			break
		}
		if days[i+curDay] == 0 {
			penalty += 40
		}
	}
	return penalty
}

// GetConsecutivePenalty checks if solution violates consecutive working/offs penalty and returns penalty
func GetConsecutivePenalty(consecutive []int, min int, max int) int {
	penalty := 0
	for _, w := range consecutive {
		if w > max {
			penalty += 20
		}
		if w < min {
			penalty += 20
		}
	}
	return penalty
}

// EvalDays evaluates day rules
func EvalDays(days []int, blocks [][]int, solution []int, minOff int, maxOff int, minWorking int, maxWorking int) []int {
	consecutiveWorkings, consecutiveOffs, totalWorking, wrongDayOffPenalty := GetDaysStats(days, blocks, solution)

	penalty := 0

	consecutiveWorkingsPenalty := GetConsecutivePenalty(consecutiveWorkings, minWorking, maxWorking)
	consecutiveOffPenalty := GetConsecutivePenalty(consecutiveOffs, minOff, maxOff)

	penalty = wrongDayOffPenalty + consecutiveWorkingsPenalty + consecutiveOffPenalty

	return []int{penalty, totalWorking}
}

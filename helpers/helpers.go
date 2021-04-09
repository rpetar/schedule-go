package helpers

import (
	"github.com/Tymeshift/data-science-test/schedule"
)

// ProblemParams is a struct with all relevant params for the problem
type ProblemParams struct {
	NumDays    int
	MaxWorking int
	MinWorking int
	MaxOff     int
	MinOff     int
	DaysOff    []int
}


var ParamSets = []ProblemParams{{
	NumDays:    28,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{5, 6},
}, {
	NumDays:    28,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{2, 5},
}, {
	NumDays:    112,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{5, 6},
}, {
	NumDays:    112,
	MaxWorking: 4,
	MinWorking: 1,
	MaxOff:     4,
	MinOff:     2,
	DaysOff:    []int{},
}, {
	NumDays:    7,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{5, 6},
}}

func isInSlice(a int, b []int) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}
	return false
}


func GetSearchParams(params ProblemParams) (schedule.Evaluate, schedule.FindNeighborhood, schedule.BlocksData) {
	// template for mandatory pre-defined day-offs
	var workingDays []int
	for i := 0; i < params.NumDays; i++ {
		if i != 0 && isInSlice(i%7, params.DaysOff) {
			// pre-defined day-off
			workingDays = append(workingDays, 0)
		} else {
			// each day have a different value associated with it, depending on the demand.
			workingDays = append(workingDays, 1)
		}
	}

	blocks := schedule.CreateBlocks(params.MinOff, params.MaxOff, params.MinWorking, params.MaxWorking)
	evaluate := schedule.EvaluateFactory(params.NumDays, workingDays, blocks.Weights, blocks.Blocks, params.MinOff, params.MaxOff, params.MinWorking, params.MaxWorking)
	findNeighborhood := schedule.FindNeighborhoodFactory(len(blocks.Weights))
	return evaluate, findNeighborhood, blocks
}

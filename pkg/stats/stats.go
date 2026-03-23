package stats

import "math"

type TargetResult struct {
	TargetVPIP    float64
	FoldsNeeded   float64
	NewTotalHands float64
	ResultingPFR  float64
}

// CalculateDrops computes the required folds to reach a specified target VPIP
func CalculateDrops(hands, vpip, pfr, targetLimit float64) []TargetResult {
	var results []TargetResult
	
	// if they are already below the target, return nothing
	if vpip <= targetLimit {
		return results
	}

	vpipHands := (hands * vpip) / 100.0
	pfrHands := (hands * pfr) / 100.0

	startTarget := math.Floor(vpip)
	if startTarget == vpip {
		startTarget--
	}

	for target := startTarget; target >= targetLimit; target-- {
		newTotalHands := (vpipHands * 100.0) / target
		foldsNeeded := newTotalHands - hands
		newPFR := (pfrHands / newTotalHands) * 100.0

		results = append(results, TargetResult{
			TargetVPIP:    target,
			FoldsNeeded:   math.Ceil(foldsNeeded),
			NewTotalHands: math.Ceil(newTotalHands),
			ResultingPFR:  newPFR,
		})
	}
	return results
}

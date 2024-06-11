package semantic_router

import "gonum.org/v1/gonum/floats"

// NormalizeScores normalizes the similarity scores to a 0-1 range.
//
// The function takes a slice of float64 values representing the
// similarity scores and returns a slice of float64 values representing
// the normalized similarity scores.
func NormalizeScores(sim []float64) []float64 {
	minVal := floats.Min(sim)
	maxVal := floats.Max(sim)
	normalized := make([]float64, len(sim))
	for i := 0; i < len(sim); i++ {
		if maxVal == minVal {
			normalized[i] = 0 // Avoid division by zero if all values are the same
		} else {
			normalized[i] = (sim[i] - minVal) / (maxVal - minVal)
		}
	}
	return normalized
}

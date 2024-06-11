package semantic_router

import (
	"testing"
)

// TestNormalizeScores tests the NormalizeScores function
// to ensure that the scores are normalized correctly.
// It does this testing by comparing the normalized scores
// with the expected values.
func TestNormalizeScores(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected []float64
	}{
		{
			input:    []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: []float64{0.0, 0.25, 0.5, 0.75, 1.0},
		},
		{
			input:    []float64{5.0, 5.0, 5.0, 5.0},
			expected: []float64{0.0, 0.0, 0.0, 0.0},
		},
		{
			input: []float64{2.0, 8.0, 4.0, 6.0},
			expected: []float64{
				0.0,
				1.0,
				0.3333333333333333,
				0.6666666666666666,
			},
		},
		{
			input:    []float64{0.0, 0.5, 1.0, 1.5, 2.0},
			expected: []float64{0.0, 0.25, 0.5, 0.75, 1.0},
		},
		{input: []float64{-1.0, 0.0, 1.0}, expected: []float64{0.0, 0.5, 1.0}},
	}
	for _, tc := range testCases {
		result := NormalizeScores(tc.input)
		for i, v := range result {
			if v != tc.expected[i] {
				t.Errorf(
					"For input %v, expected %v but got %v",
					tc.input,
					tc.expected,
					result,
				)
				break
			}
		}
	}
}

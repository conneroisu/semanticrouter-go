package semantic_router

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

// TestSimilarityMatrix tests the similarityMatrix function
func TestSimilarityMatrix(t *testing.T) {
	tests := []struct {
		name   string
		xq     *mat.VecDense
		index  *mat.Dense
		expect *mat.VecDense
	}{
		{
			name:  "Basic test",
			xq:    mat.NewVecDense(3, []float64{1, 2, 3}),
			index: mat.NewDense(4, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			expect: mat.NewVecDense(4, []float64{
				1.0000000000000002,
				0.9746318461970761,
				0.9594119455666704,
				0.9512583076673059,
			}),
		},
		{
			name:  "Zero vector query",
			xq:    mat.NewVecDense(3, []float64{0, 0, 0}),
			index: mat.NewDense(4, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			expect: mat.NewVecDense(4, []float64{
				0,
				0,
				0,
				0,
			}),
		},
		{
			name:  "Single element index",
			xq:    mat.NewVecDense(3, []float64{1, 2, 3}),
			index: mat.NewDense(1, 3, []float64{1, 2, 3}),
			expect: mat.NewVecDense(1, []float64{
				1,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SimilarityMatrix(tt.xq, tt.index)
			if !mat.EqualApprox(result, tt.expect, 1e-6) {
				t.Errorf("got %v, want %v", result, tt.expect)
			}
		})
	}
}

// TestTopScores tests the topScores function
func TestTopScores(t *testing.T) {
	tests := []struct {
		name        string
		sim         *mat.VecDense
		topK        int
		wantScores  []float64
		wantIndices []int
	}{
		{
			name:        "Basic test",
			sim:         mat.NewVecDense(4, []float64{1, 0.5, 0.2, 0.8}),
			topK:        2,
			wantScores:  []float64{1, 0.8},
			wantIndices: []int{0, 3},
		},
		{
			name:        "Empty index",
			sim:         mat.NewVecDense(4, []float64{1, 0.5, 0.2, 0.8}),
			topK:        2,
			wantScores:  []float64{1, 0.8},
			wantIndices: []int{0, 3},
		},
		{
			name:        "Empty query",
			sim:         mat.NewVecDense(4, []float64{1, 0.5, 0.2, 0.8}),
			topK:        2,
			wantScores:  []float64{1, 0.8},
			wantIndices: []int{0, 3},
		},
		{
			name:        "Empty query and index",
			sim:         mat.NewVecDense(4, []float64{1, 0.5, 0.2, 0.8}),
			topK:        2,
			wantScores:  []float64{1, 0.8},
			wantIndices: []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScores, gotIndices := TopScores(tt.sim, tt.topK)
			if !reflect.DeepEqual(gotScores, tt.wantScores) {
				t.Errorf("got scores %v, want %v", gotScores, tt.wantScores)
			}
			if !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("got indices %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}

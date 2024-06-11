package semantic_router

import (
	"math"
	"sort"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// SimilarityMatrix computes the similarity scores between a query vector and a set of vectors.
func SimilarityMatrix(xq, index mat.Matrix) *mat.VecDense {
	rows, cols := index.Dims()
	xqVec, ok := xq.(*mat.VecDense)
	if !ok {
		xqVec = mat.NewVecDense(cols, nil)
		for i := 0; i < cols; i++ {
			xqVec.SetVec(i, xq.At(i, 0))
		}
	}
	xqNorm := mat.Norm(xqVec, 2)
	indexNorm := make([]float64, rows)
	for i := 0; i < rows; i++ {
		rowVec := mat.Row(nil, i, index)
		indexNorm[i] = floats.Norm(rowVec, 2)
	}
	sim := make([]float64, rows)
	for i := 0; i < rows; i++ {
		rowVec := mat.Row(nil, i, index)
		dot := floats.Dot(rowVec, xqVec.RawVector().Data)
		sim[i] = dot / (indexNorm[i] * xqNorm)
	}
	// if a vecot is NaN, it will be replaced by 0
	for i, v := range sim {
		if math.IsNaN(v) {
			sim[i] = 0
		}
	}
	return mat.NewVecDense(rows, sim)
}

// TopScores returns the scores and indices of the top k scores from the similarity matrix.
func TopScores(sim *mat.VecDense, topK int) ([]float64, []int) {
	s := sim.RawVector().Data
	if topK > len(s) {
		topK = len(s)
	}
	type scoreIndex struct {
		score float64
		index int
	}
	si := make([]scoreIndex, len(s))
	for i, score := range s {
		si[i] = scoreIndex{score, i}
	}
	sort.Slice(si, func(i, j int) bool {
		return si[i].score > si[j].score
	})
	topScores := make([]float64, topK)
	topIndices := make([]int, topK)
	for i := 0; i < topK; i++ {
		topScores[i] = si[i].score
		topIndices[i] = si[i].index
	}
	return topScores, topIndices
}

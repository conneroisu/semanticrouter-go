package semantic_router

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// SimilarityMatrix computes the similarity scores between a query vector and a set of vectors.
//
// The similarity score is the dot product of the query vector and the index vector divided by
// the product of the query vector norm and the index vector norm.
//
// The similarity matrix returned is a matrix where each element is the similarity score between the query
// vector and the corresponding index vector.
func SimilarityMatrix(xq, index *mat.VecDense) float64 {
	xqNorm := mat.Norm(xq, 2)
	indexNorm := mat.Norm(index, 2)
	dot := floats.Dot(xq.RawVector().Data, index.RawVector().Data)
	return dot / (xqNorm * indexNorm)
}

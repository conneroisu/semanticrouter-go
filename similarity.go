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
	// normalize the query vector
	xqNorm := mat.Norm(xq, 2)
	// normalize the index vector
	indexNorm := mat.Norm(index, 2)
	// perform dot product between query vector and index vector
	dot := floats.Dot(xq.RawVector().Data, index.RawVector().Data)
	// return the similarity score (dot product) divided by the product of the query vector norm and the index vector norm
	return dot / (xqNorm * indexNorm)
}

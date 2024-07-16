package semanticrouter

import (
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// SimilarityDotMatrix computes the similarity scores between a query vector and a set of vectors.
//
// The similarity score is the dot product of the query vector and the index vector divided by
// the product of the query vector norm and the index vector norm.
//
// The similarity matrix returned is a matrix where each element is the similarity score between the query
// vector and the corresponding index vector.
func SimilarityDotMatrix(xq, index *mat.VecDense) float64 {
	// normalize the query vector
	xqNorm := mat.Norm(xq, 2)
	// normalize the index vector
	indexNorm := mat.Norm(index, 2)
	// perform dot product between query vector and index vector
	dot := floats.Dot(xq.RawVector().Data, index.RawVector().Data)
	// return the similarity score (dot product) divided by the product of the query vector norm and the index vector norm
	return dot / (xqNorm * indexNorm)
}

// EuclideanDistance calculates the Euclidean distance between two vectors.
//
// The euclidean distance is the square root of the sum of the squared differences between corresponding elements in the two vectors.
//
// The function takes two vectors as input and returns the Euclidean distance between them.
func EuclideanDistance(xq, index *mat.VecDense) float64 {
	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	return mat.Norm(diff, 2)
}

// ManhattanDistance calculates the Manhattan distance between two vectors.
//
// The manhattan distance is the sum of the absolute differences between corresponding elements in the two vectors.
//
// The function takes two vectors as input and returns the Manhattan distance between them.
func ManhattanDistance(xq, index *mat.VecDense) float64 {
	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	sum := 0.0
	for i := 0; i < diff.Len(); i++ {
		sum += math.Abs(diff.AtVec(i))
	}
	return mat.Norm(diff, 1)
}

// JaccardSimilarity calculates the Jaccard similarity between two vectors.
//
// The Jaccard similarity is the size of the intersection of the two sets divided by the size of the union of the two sets.
//
// The function takes two vectors as input and returns the Jaccard distance between them.
func JaccardSimilarity(xq, index *mat.VecDense) float64 {
	minSum := 0.0
	maxSum := 0.0
	for i := 0; i < xq.Len(); i++ {
		minSum += math.Min(xq.AtVec(i), index.AtVec(i))
		maxSum += math.Max(xq.AtVec(i), index.AtVec(i))
	}
	return minSum / maxSum
}

// PearsonCorrelation calculates the Pearson correlation between two vectors.
//
// The Pearson correlation is a measure of the linear relationship between two variables.
// It ranges from -1 to 1, where -1 indicates a perfect negative correlation, 0 indicates no correlation, and 1 indicates a perfect positive correlation.
//
// The function takes two vectors as input and returns the Pearson correlation between them.
func PearsonCorrelation(xq, index *mat.VecDense) float64 {
	meanXq := floats.Sum(xq.RawVector().Data) / float64(xq.Len())
	meanIndex := floats.Sum(index.RawVector().Data) / float64(index.Len())

	numerator := 0.0
	varSumXq := 0.0
	varSumIndex := 0.0

	for i := 0; i < xq.Len(); i++ {
		diffXq := xq.AtVec(i) - meanXq
		diffIndex := index.AtVec(i) - meanIndex
		numerator += diffXq * diffIndex
		varSumXq += diffXq * diffXq
		varSumIndex += diffIndex * diffIndex
	}

	return numerator / (math.Sqrt(varSumXq) * math.Sqrt(varSumIndex))
}

// HammingDistance calculates the Hamming distance between two vectors.
//
// The Hamming distance is the number of positions at which the corresponding bits are different.
//
// The function takes two vectors as input and returns the Hamming distance between them.
func HammingDistance(xq, index *mat.VecDense) float64 {
	if xq.Len() != index.Len() {
		panic("Vectors must be the same length")
	}

	count := 0.0
	for i := 0; i < xq.Len(); i++ {
		if xq.AtVec(i) != index.AtVec(i) {
			count++
		}
	}

	return count
}

// MinkowskiDistance calculates the Minkowski distance between two vectors.
func MinkowskiDistance(xq, index *mat.VecDense, p float64) float64 {
	if p <= 0 {
		panic("Order p must be greater than 0")
	}

	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	sum := 0.0
	for i := 0; i < diff.Len(); i++ {
		sum += math.Pow(math.Abs(diff.AtVec(i)), p)
	}

	return math.Pow(sum, 1/p)
}

package semanticrouter

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// NormalizeScores normalizes the similarity scores to a 0-1 range.
// The function takes a slice of float64 values representing the similarity
// scores.
//
// The function takes a slice of float64 values representing the
// similarity scores and returns a slice of float64 values representing
// the normalized similarity scores.
func NormalizeScores(sim []float64) []float64 {
	minimum := floats.Min(sim)
	maximum := floats.Max(sim)
	normalized := make([]float64, len(sim))
	for i := 0; i < len(sim); i++ {
		if maximum == minimum {
			// Avoid division by zero if all values are the same
			normalized[i] = 0
		} else {
			normalized[i] = (sim[i] - minimum) / (maximum - minimum)
		}
	}
	return normalized
}

// WithSimilarityDotMatrix sets the similarity function to use with a coefficient.
func WithSimilarityDotMatrix(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeff = append(r.biFuncCoeff, biFuncCoefficient{
			handler:     SimilarityDotMatrix,
			coefficient: coefficient,
		})
	}
}

// WithEuclideanDistance sets the EuclideanDistance function with a coefficient.
func WithEuclideanDistance(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeff = append(r.biFuncCoeff, biFuncCoefficient{
			handler:     EuclideanDistance,
			coefficient: coefficient,
		})
	}
}

// WithManhattanDistance sets the ManhattanDistance function with a coefficient.
func WithManhattanDistance(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeff = append(r.biFuncCoeff, biFuncCoefficient{
			handler:     ManhattanDistance,
			coefficient: coefficient,
		})
	}
}

// WithJaccardSimilarity sets the JaccardSimilarity function with a coefficient.
func WithJaccardSimilarity(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeff = append(r.biFuncCoeff, biFuncCoefficient{
			handler:     JaccardSimilarity,
			coefficient: coefficient,
		})
	}
}

// WithPearsonCorrelation sets the PearsonCorrelation function with a coefficient.
func WithPearsonCorrelation(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeff = append(r.biFuncCoeff, biFuncCoefficient{
			handler:     PearsonCorrelation,
			coefficient: coefficient,
		})
	}
}

// SimilarityDotMatrix computes the similarity scores between a query vector and
// a set of vectors.
//
// The similarity score is the dot product of the query vector and the index
// vector divided by the product of the query vector norm and the index vector
// norm.
//
// $$a \cdot b=\sum_{i=1}^{n} a_{i} b_{i}$$
//
// The similarity matrix returned is a matrix where each element is reduced to
// the similarity score between the query vector and the corresponding index
// vector.
func SimilarityDotMatrix(xq, index *mat.VecDense) (float64, error) {
	// normalize the query vector
	xqNorm := mat.Norm(xq, 2)
	// normalize the index vector
	indexNorm := mat.Norm(index, 2)
	// perform dot product between query vector and index vector
	dot := floats.Dot(xq.RawVector().Data, index.RawVector().Data)
	// return the similarity score (dot product) divided by the product of
	// the query vector norm and the index vector norm
	return dot / (xqNorm * indexNorm), nil
}

// EuclideanDistance calculates the Euclidean distance between two vectors.
//
// The euclidean distance is the square root of the sum of the squared
// differences between corresponding elements in the two vectors.
//
// $$d(x, y) = \sqrt{\sum_{i=1}^{n}(x_i - y_i)^2}$$
//
// The function takes two vectors as input and returns the Euclidean distance
// between them.
func EuclideanDistance(xq, index *mat.VecDense) (float64, error) {
	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	return mat.Norm(diff, 2), nil
}

// ManhattanDistance calculates the Manhattan distance between two vectors.
//
// The manhattan distance is the sum of the absolute differences between
// corresponding elements in the two vectors.
//
// $$d(x, y) = |x_1 - y_1| + |x_2 - y_2| + ... + |x_n - y_n|$$
//
// The function takes two vectors as input and returns the Manhattan distance between them.
func ManhattanDistance(xq, index *mat.VecDense) (float64, error) {
	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	sum := 0.0
	for i := 0; i < diff.Len(); i++ {
		sum += math.Abs(diff.AtVec(i))
	}
	return mat.Norm(diff, 1), nil
}

// JaccardSimilarity calculates the Jaccard similarity between two vectors.
//
// The Jaccard similarity is the size of the intersection of the two sets divided by the size of the union of the two sets.
//
// $$J(A, B)=\frac{|A \cap B|}{|A \cup B|}$$
//
// The function takes two vectors as input and returns the Jaccard distance between them.
func JaccardSimilarity(xq, index *mat.VecDense) (float64, error) {
	minSum := 0.0
	maxSum := 0.0
	for i := 0; i < xq.Len(); i++ {
		minSum += math.Min(xq.AtVec(i), index.AtVec(i))
		maxSum += math.Max(xq.AtVec(i), index.AtVec(i))
	}
	return minSum / maxSum, nil
}

// PearsonCorrelation calculates the Pearson correlation between two vectors.
//
// The Pearson correlation is a measure of the linear relationship between two
// variables.
//
// It ranges from -1 to 1, where -1 indicates a perfect negative correlation, 0
// indicates no correlation, and 1 indicates a perfect positive correlation.
//
// $$r=\frac{\sum\left(x_{i}-\bar{x}\right)\left(y_{i}-\bar{y}\right)}{\sqrt{\sum\left(x_{i}-\bar{x}\right)^{2} \sum\left(y_{i}-\bar{y}\right)^{2}}}$$
//
// $r$ correlation coefficient
//
// $x_{i}$ values of the $$x$$-variable in a sample
//
// $\bar{x}$ is the mean of the values of the $$x$$-variable
//
// $y_{i}$ are the values of the $$y$$-variable in a sample
//
// $\bar{y}$ is the mean of the values of the $$y$$-variable
//
// The function takes two vectors as input and returns the Pearson correlation
// between them.
func PearsonCorrelation(xq, index *mat.VecDense) (float64, error) {
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

	return numerator / (math.Sqrt(varSumXq) * math.Sqrt(varSumIndex)), nil
}

// HammingDistance calculates the Hamming distance between two vectors.
//
// The Hamming distance is the number of positions at which the corresponding
// bits are different.
//
// $$d(x, y)=\frac{1}{n} \sum_{n=1}^{n=n}\left|x_{i}-y_{i}\right|$$
//
// The function takes two vectors as input and returns the Hamming distance between them.
func HammingDistance(xq, index *mat.VecDense) (float64, error) {
	if xq.Len() != index.Len() {
		return 0, fmt.Errorf("Vectors must be the same length")
	}

	count := 0.0
	for i := 0; i < xq.Len(); i++ {
		if xq.AtVec(i) != index.AtVec(i) {
			count++
		}
	}

	return count, nil
}

// MinkowskiDistance calculates the Minkowski distance between two vectors.
//
// The Minkowski distance is the sum of the absolute differences between
// corresponding elements in the two vectors raised to the power of p.
//
// $$ d(x, y) = \sum_{i=1}^{n} |x_i - y_i|^p $$
//
// where n is the length of the vectors.
func MinkowskiDistance(xq, index *mat.VecDense, p float64) (float64, error) {
	if p <= 0 {
		panic("Order p must be greater than 0")
	}

	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	sum := 0.0
	for i := 0; i < diff.Len(); i++ {
		sum += math.Pow(math.Abs(diff.AtVec(i)), p)
	}

	return math.Pow(sum, 1/p), nil
}

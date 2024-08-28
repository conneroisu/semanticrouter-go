package semanticrouter

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// embedding is the embedding of some text, speech, or other data (images, videos, etc.).
type embedding []float64

// Utterance represents a utterance in the semantic router.
type Utterance struct {
	// ID is the ID of the utterance.
	ID int
	// Utterance is the text of the utterance.
	Utterance string
	// Embed is the embedding of the utterance. It is a vector of floats.
	Embed embedding
}

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

// WithSimilarityDotMatrix sets the similarity function to use with a
// coefficient.
//
// $$a \cdot b=\sum_{i=1}^{n} a_{i} b_{i}$$
//
// It adds the similarity dot matrix to the comparision functions with the given
// coefficient.
func WithSimilarityDotMatrix(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeffs = append(r.biFuncCoeffs, biFuncCoefficient{
			handler:     similarityDotMatrix,
			coefficient: coefficient,
		})
	}
}

// WithEuclideanDistance sets the EuclideanDistance function with a coefficient.
//
// $$d(x, y) = \sqrt{\sum_{i=1}^{n}(x_i - y_i)^2}$$
func WithEuclideanDistance(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeffs = append(r.biFuncCoeffs, biFuncCoefficient{
			handler:     euclideanDistance,
			coefficient: coefficient,
		})
	}
}

// WithManhattanDistance sets the ManhattanDistance function with a coefficient.
//
// $$d(x, y) = |x_1 - y_1| + |x_2 - y_2| + ... + |x_n - y_n|$$
//
// It adds the manhatten distance to the comparision functions with the given
// coefficient.
func WithManhattanDistance(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeffs = append(r.biFuncCoeffs, biFuncCoefficient{
			handler:     manhattanDistance,
			coefficient: coefficient,
		})
	}
}

// WithJaccardSimilarity sets the JaccardSimilarity function with a coefficient.
//
// $$J(A, B)=\frac{|A \cap B|}{|A \cup B|}$$
//
// It adds the jaccard similarity to the comparision functions with the given
// coefficient.
func WithJaccardSimilarity(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeffs = append(r.biFuncCoeffs, biFuncCoefficient{
			handler:     jaccardSimilarity,
			coefficient: coefficient,
		})
	}
}

// WithPearsonCorrelation sets the PearsonCorrelation function with a
// coefficient.
//
// $$r=\frac{\sum\left(x_{i}-\bar{x}\right)\left(y_{i}-\bar{y}\right)}{\sqrt{\sum\left(x_{i}-\bar{x}\right)^{2} \sum\left(y_{i}-\bar{y}\right)^{2}}}$$
//
// It adds the pearson correlation to the comparision functions with the given
// coefficient.
func WithPearsonCorrelation(coefficient float64) Option {
	return func(r *Router) {
		r.biFuncCoeffs = append(r.biFuncCoeffs, biFuncCoefficient{
			handler:     pearsonCorrelation,
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
func similarityDotMatrix(xq, index *mat.VecDense) (float64, error) {
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
func euclideanDistance(xq, index *mat.VecDense) (float64, error) {
	diff := mat.NewVecDense(xq.Len(), nil)
	diff.SubVec(xq, index)
	return mat.Norm(diff, 2), nil
}

// manhattanDistance calculates the Manhattan distance between two vectors.
//
// The manhattan distance is the sum of the absolute differences between
// corresponding elements in the two vectors.
//
// $$d(x, y) = |x_1 - y_1| + |x_2 - y_2| + ... + |x_n - y_n|$$
//
// The function takes two vectors as input and returns the Manhattan distance between them.
func manhattanDistance(xq, index *mat.VecDense) (float64, error) {
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
func jaccardSimilarity(xq, index *mat.VecDense) (float64, error) {
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
// The function takes two vectors as input and returns the Pearson correlation
// between them.
func pearsonCorrelation(xq, index *mat.VecDense) (float64, error) {
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
func hammingDistance(xq, index *mat.VecDense) (float64, error) {
	if xq.Len() != index.Len() {
		return 0, fmt.Errorf("vectors must be the same length (are you mix mashing encoding models?)")
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
// $$d(x, y) = \sum_{i=1}^{n} |x_i - y_i|^p$$
//
// where n is the length of the vectors.
func minkowskiDistance(xq, index *mat.VecDense, p float64) (float64, error) {
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

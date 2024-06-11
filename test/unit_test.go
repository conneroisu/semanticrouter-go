package test

import (
	"fmt"
	"testing"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"gonum.org/v1/gonum/mat"
)

func TestSimilarityMatrix(t *testing.T) {
	// Example usage
	// Define query vector xq and index matrix
	xq := mat.NewVecDense(3, []float64{1, 2, 3})
	index := mat.NewDense(4, 3, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		10, 11, 12,
	})

	// Compute similarity matrix
	sim := semantic_router.SimilarityMatrix(xq, index)
	fmt.Println("Similarity scores:", sim.RawVector().Data)

	// Get top scores
	scores, indices := semantic_router.TopScores(sim, 2)
	fmt.Println("Top scores:", scores)
	fmt.Println("Indices of top scores:", indices)

}

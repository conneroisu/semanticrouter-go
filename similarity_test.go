package semanticrouter

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// Helper function to create a VecDense from a slice
func createVecDense(data []float64) *mat.VecDense {
	return mat.NewVecDense(len(data), data)
}

// Test case struct to hold query, index and expected similarity value
type testCase struct {
	queryVec    []float64
	indexVec    []float64
	expectedSim float64
}

// TestSimilarityMatrix tests the SimilarityMatrix function
func TestSimilarityMatrix(t *testing.T) {
	tests := []testCase{
		{
			queryVec: []float64{
				0.6378429023635818,
				0.6891080666053131,
				0.6242938529238299,
				0.44797618387108773,
				0.28283025927535843,
				0.7999294372242242,
				0.8289827972810941,
			},
			indexVec: []float64{
				0.7281087474470895,
				0.1911238756245191,
				0.5368592300231692,
				0.2210151126530714,
				0.5113255269750295,
				0.260703208744612,
				0.7499797968916341,
			},
			expectedSim: 0.8819,
		},
		{
			queryVec: []float64{
				0.3213532863532023,
				0.01524713642631278,
				0.5640214803262418,
				0.7471951467346923,
			},
			indexVec: []float64{
				0.14265224091380074,
				0.5373162226984148,
				0.7329499385535614,
				0.11489132191465051,
			},
			expectedSim: 0.6029,
		},
		{
			queryVec: []float64{
				0.3679798861563019,
				0.006742820180837044,
				0.3370298486376625,
				0.9643456114675119,
				0.9445383919088846,
				0.9734395554155766,
				0.9577395250396427,
				0.3953503523165984,
			},
			indexVec: []float64{
				0.42518640372707744,
				0.5552344499075671,
				0.5756751288905184,
				0.49370824448686684,
				0.121299316718655,
				0.4967735216023797,
				0.2780113858159054,
				0.9210994797770689,
			},
			expectedSim: 0.6783,
		},
		{
			queryVec: []float64{
				0.8845173513113452,
				0.759379579636257,
				0.7739941929392022,
				0.43620853139649546,
			},
			indexVec: []float64{
				0.41753838520979547,
				0.17957609228891147,
				0.9254647571107574,
				0.4308401190361257,
			},
			expectedSim: 0.8608,
		},
		{
			queryVec: []float64{
				0.623464705485095,
				0.6391318524794283,
				0.5521209781501045,
				0.7600363967129513,
			},
			indexVec: []float64{
				0.17958471999594827,
				0.8172266806505036,
				0.31984582091877944,
				0.8880934772457741,
			},
			expectedSim: 0.9089,
		},
		{
			queryVec:    []float64{0.05336304768380549},
			indexVec:    []float64{0.7239604234641187},
			expectedSim: 1.0000,
		},
		{
			queryVec: []float64{
				0.9259701759307833,
				0.4369527176226245,
				0.0009199576941947202,
				0.10025644542794729,
			},
			indexVec: []float64{
				0.24059696635437425,
				0.023210885389478467,
				0.11345058443817552,
				0.3520946084651303,
			},
			expectedSim: 0.5902,
		},
		{
			queryVec: []float64{
				0.7426830477371793,
				0.38641704712181063,
				0.8201992083605694,
				0.0928964204029426,
				0.06200215665820996,
			},
			indexVec: []float64{
				0.19359892971119344,
				0.29249943767771514,
				0.22711180425384944,
				0.6436326698081482,
				0.6220624241769184,
			},
			expectedSim: 0.4656,
		},
		{
			queryVec: []float64{
				0.6591404347120998,
				0.9940819471256389,
				0.1901385177574703,
				0.8717107048444336,
				0.8297715360829513,
			},
			indexVec: []float64{
				0.1838956335011302,
				0.09479084716976331,
				0.450196905643527,
				0.8206225692172371,
				0.7116401506909583,
			},
			expectedSim: 0.7894,
		},
		{
			queryVec: []float64{
				0.44694905646326777,
				0.20039786752144578,
				0.5473983349535733,
			},
			indexVec: []float64{
				0.4628576782809643,
				0.5124668827802493,
				0.40255295053932205,
			},
			expectedSim: 0.9026,
		},
	}
	for _, tc := range tests {
		t.Run(
			fmt.Sprintf(
				"TestSimilarityMatrix(%v, %v)",
				tc.queryVec,
				tc.indexVec,
			),
			func(t *testing.T) {
				a := assert.New(t)
				tc := tc
				t.Parallel()
				query := createVecDense(tc.queryVec)
				index := createVecDense(tc.indexVec)
				similarity, err := similarityDotMatrix(
					query,
					index,
				)
				a.NoError(err)

				expected := []float64{tc.expectedSim}
				actual := []float64{similarity}

				if !floats.EqualApprox(expected, actual, 0.0001) {
					t.Errorf(
						"SimilarityMatrix(%v, %v) = %v; want %v",
						tc.queryVec,
						tc.indexVec,
						similarity,
						tc.expectedSim,
					)
				}
			},
		)
	}
}

// TestSimilarityDotMatrix tests the SimilarityDotMatrix function
func TestSimilarityDotMatrix(t *testing.T) {
	tests := []struct {
		xq, index []float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 0.9746318461970762},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 0},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 1},
	}

	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := similarityDotMatrix(xq, index)
		a.NoError(err)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf(
				"SimilarityDotMatrix(%v, %v) = %v; want %v",
				tt.xq,
				tt.index,
				got,
				tt.want,
			)
		}
	}
}

// TestEuclideanDistance tests the EuclideanDistance function
func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		xq, index []float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 5.196152422706632},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 1.7320508075688772},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 0},
	}

	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := euclideanDistance(xq, index)
		a.NoError(err)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf(
				"EuclideanDistance(%v, %v) = %v; want %v",
				tt.xq,
				tt.index,
				got,
				tt.want,
			)
		}
	}
}

// TestManhattanDistance tests the ManhattanDistance function
func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		xq, index []float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 9},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 3},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 0},
	}

	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := manhattanDistance(xq, index)
		a.NoError(err)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf(
				"ManhattanDistance(%v, %v) = %v; want %v",
				tt.xq,
				tt.index,
				got,
				tt.want,
			)
		}
	}
}

// TestJaccardSimilarity tests the JaccardSimilarity function
func TestJaccardSimilarity(t *testing.T) {
	tests := []struct {
		xq, index []float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 0.4},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 0},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 1},
	}

	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := jaccardSimilarity(xq, index)
		a.NoError(err)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf(
				"JaccardSimilarity(%v, %v) = %v; want %v",
				tt.xq,
				tt.index,
				got,
				tt.want,
			)
		}
	}
}

// TestPearsonCorrelation tests the PearsonCorrelation function
func TestPearsonCorrelation(t *testing.T) {
	tests := []struct {
		xq, index []float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 1},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 0},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, math.NaN()},
	}
	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := pearsonCorrelation(xq, index)
		a.NoError(err)
		if math.IsNaN(tt.want) {
			if !math.IsNaN(got) {
				t.Errorf(
					"PearsonCorrelation(%v, %v) = %v; want %v",
					tt.xq,
					tt.index,
					got,
					tt.want,
				)
			}
		} else if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("PearsonCorrelation(%v, %v) = %v; want %v", tt.xq, tt.index, got, tt.want)
		}
	}
}

// TestHammingDistance tests the HammingDistance function
func TestHammingDistance(t *testing.T) {
	tests := []struct {
		xq, index []float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 3},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 3},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 0},
	}

	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := hammingDistance(xq, index)
		a.NoError(err)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf(
				"HammingDistance(%v, %v) = %v; want %v",
				tt.xq,
				tt.index,
				got,
				tt.want,
			)
		}
	}
}

// TestMinkowskiDistance tests the MinkowskiDistance function
func TestMinkowskiDistance(t *testing.T) {
	tests := []struct {
		xq, index []float64
		p         float64
		want      float64
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, 3, 4.326748710922225},
		{[]float64{0, 0, 0}, []float64{1, 1, 1}, 2, 1.7320508075688772},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 1, 0},
	}

	for _, tt := range tests {
		a := assert.New(t)
		xq := mat.NewVecDense(len(tt.xq), tt.xq)
		index := mat.NewVecDense(len(tt.index), tt.index)
		got, err := minkowskiDistance(xq, index, tt.p)
		a.NoError(err)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf(
				"MinkowskiDistance(%v, %v, %v) = %v; want %v",
				tt.xq,
				tt.index,
				tt.p,
				got,
				tt.want,
			)
		}
	}
}

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

package semanticrouter

import (
	"fmt"
	"testing"

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
			queryVec:    []float64{0.6378429023635818, 0.6891080666053131, 0.6242938529238299, 0.44797618387108773, 0.28283025927535843, 0.7999294372242242, 0.8289827972810941},
			indexVec:    []float64{0.7281087474470895, 0.1911238756245191, 0.5368592300231692, 0.2210151126530714, 0.5113255269750295, 0.260703208744612, 0.7499797968916341},
			expectedSim: 0.8819,
		},
		{
			queryVec:    []float64{0.3213532863532023, 0.01524713642631278, 0.5640214803262418, 0.7471951467346923},
			indexVec:    []float64{0.14265224091380074, 0.5373162226984148, 0.7329499385535614, 0.11489132191465051},
			expectedSim: 0.6029,
		},
		{
			queryVec:    []float64{0.3679798861563019, 0.006742820180837044, 0.3370298486376625, 0.9643456114675119, 0.9445383919088846, 0.9734395554155766, 0.9577395250396427, 0.3953503523165984},
			indexVec:    []float64{0.42518640372707744, 0.5552344499075671, 0.5756751288905184, 0.49370824448686684, 0.121299316718655, 0.4967735216023797, 0.2780113858159054, 0.9210994797770689},
			expectedSim: 0.6783,
		},
		{
			queryVec:    []float64{0.8845173513113452, 0.759379579636257, 0.7739941929392022, 0.43620853139649546},
			indexVec:    []float64{0.41753838520979547, 0.17957609228891147, 0.9254647571107574, 0.4308401190361257},
			expectedSim: 0.8608,
		},
		{
			queryVec:    []float64{0.623464705485095, 0.6391318524794283, 0.5521209781501045, 0.7600363967129513},
			indexVec:    []float64{0.17958471999594827, 0.8172266806505036, 0.31984582091877944, 0.8880934772457741},
			expectedSim: 0.9089,
		},
		{
			queryVec:    []float64{0.05336304768380549},
			indexVec:    []float64{0.7239604234641187},
			expectedSim: 1.0000,
		},
		{
			queryVec:    []float64{0.9259701759307833, 0.4369527176226245, 0.0009199576941947202, 0.10025644542794729},
			indexVec:    []float64{0.24059696635437425, 0.023210885389478467, 0.11345058443817552, 0.3520946084651303},
			expectedSim: 0.5902,
		},
		{
			queryVec:    []float64{0.7426830477371793, 0.38641704712181063, 0.8201992083605694, 0.0928964204029426, 0.06200215665820996},
			indexVec:    []float64{0.19359892971119344, 0.29249943767771514, 0.22711180425384944, 0.6436326698081482, 0.6220624241769184},
			expectedSim: 0.4656,
		},
		{
			queryVec:    []float64{0.6591404347120998, 0.9940819471256389, 0.1901385177574703, 0.8717107048444336, 0.8297715360829513},
			indexVec:    []float64{0.1838956335011302, 0.09479084716976331, 0.450196905643527, 0.8206225692172371, 0.7116401506909583},
			expectedSim: 0.7894,
		},
		{
			queryVec:    []float64{0.44694905646326777, 0.20039786752144578, 0.5473983349535733},
			indexVec:    []float64{0.4628576782809643, 0.5124668827802493, 0.40255295053932205},
			expectedSim: 0.9026,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestSimilarityMatrix(%v, %v)", tc.queryVec, tc.indexVec), func(t *testing.T) {
			tc := tc
			t.Parallel()
			query := createVecDense(tc.queryVec)
			index := createVecDense(tc.indexVec)
			similarity := SimilarityMatrix(query, index)

			expected := []float64{tc.expectedSim}
			actual := []float64{similarity}

			if !floats.EqualApprox(expected, actual, 0.0001) {
				t.Errorf("SimilarityMatrix(%v, %v) = %v; want %v", tc.queryVec, tc.indexVec, similarity, tc.expectedSim)
			}
		})
	}
}

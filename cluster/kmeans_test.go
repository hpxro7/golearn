package cluster

import "testing"

// Tests for k-means clustering algorithm

var normTests = []struct {
	first    []float64
	second   []float64
	expected float64
}{
	{ // zero case
		[]float64{0, 0, 0, 0},
		[]float64{0, 0, 0, 0},
		0,
	},
	{ // equal vectors case
		[]float64{9, -5, 0, 1},
		[]float64{9, -5, 0, 1},
		0,
	},
	{ // floating point testing
		[]float64{-0.1, 0.3, -0.7},
		[]float64{0.5, -0.8, 3.2},
		16.78,
	},
	{ // symmetry testing
		[]float64{1, 1, 1},
		[]float64{0.5, 0.5, 0.5},
		0.75,
	},
	{ // symmetry testing
		[]float64{0.5, 0.5, 0.5},
		[]float64{1, 1, 1},
		0.75,
	},
	{ // norm to origin case
		[]float64{12, -5, 13, 21, 2},
		[]float64{0, 0, 0, 0, 0},
		783,
	},
	{ // negatives
		[]float64{-1, -2, -3, -4, -5, -6},
		[]float64{-6, -5, -4, -3, -2, -1},
		70,
	},
	{ // scalar test
		[]float64{9},
		[]float64{12},
		9,
	},
	{
		[]float64{143.34, -0.45, 39, -99, -1491, -2.5},
		[]float64{-121.32, 81.9, 81.3, 100, 400, 0},
		3.6941039781e+6,
	},
}

func TestSquaredNorm(t *testing.T) {
	for _, test := range normTests {
		r := squaredNorm(test.first, test.second)
		if r != test.expected {
			t.Error(
				"For inputs", test.first, ",", test.second,
				", expected:", test.expected,
				", got:", r,
			)
		}
	}
}

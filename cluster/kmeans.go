package cluster

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hpxro7/golearn/base"
)

type KMeans struct {
	Clusters     int
	TrainingData *base.Instances
	iterations   int
	init         InitBy
	tolerance    float64
}

type InitBy func(*KMeans) [][]float64

var (
	KMeansPlus InitBy = func(kMeans *KMeans) [][]float64 {
		return nil
	}

	KRandomized InitBy = func(kMeans *KMeans) [][]float64 {
		data := kMeans.TrainingData
		var centroids [][]float64
		randomRows := rand.Perm(data.Rows)[:kMeans.Clusters]
		for _, row := range randomRows {
			centroids = append(centroids, data.GetRowVector(row))
		}
		return centroids
	}
)

func NewKMeans(clusters, iterations int, init InitBy, tolerance float64) *KMeans {
	cluster := &KMeans{
		Clusters:   clusters,
		init:       init,
		iterations: iterations,
		tolerance:  tolerance,
	}
	return cluster
}

func (kmeans *KMeans) Fit(data *base.Instances) {
	if data.Rows < kmeans.Clusters {
		panic("kmeans: Fewer training instances than clusters")
	}
	kmeans.TrainingData = data
	centroids := kmeans.init(kmeans)
	if len(centroids) == 0 {
		panic("kmeans: Centroid slice should be non-empty")
	}

	for i := 0; i < kmeans.iterations; i++ {
		kmeans.fitSingle(centroids)
		/*delta := 1e5
		clusterAssign := make(map[]int)
		for delta > kmeans.tolerance {

		}*/
	}
}

func (kmeans *KMeans) fitSingle(oldCentroids [][]float64) (map[int][]float64, float64) {
	labels := clusterAssign(oldCentroids, kmeans.TrainingData)
	fmt.Println(labels)
	return nil, 0
	// Centroid move
	// Compute inertia / distortion error
}

func clusterAssign(centroids [][]float64, data *base.Instances) (labels map[int][]float64) {
	labels = make(map[int][]float64, data.Rows)

	for row := 0; row < data.Rows; row++ {
		minNorm, minCentroid := squaredNorm(data.GetRowVector(row), centroids[0]), 0
		for i := 1; i < len(centroids); i++ {
			currNorm := squaredNorm(data.GetRowVector(row), centroids[i])
			if currNorm < minNorm {
				minNorm, minCentroid = currNorm, i
			}
		}
		labels[row] = centroids[minCentroid]
	}
	return
}

func squaredNorm(centroid, dataPoint []float64) float64 {
	result := 0.0
	dimensions := len(centroid)
	for i := 0; i < dimensions; i++ {
		result += math.Pow(centroid[i]-dataPoint[i], 2)
	}
	return result
}

func (kmeans *KMeans) Predict(data *base.Instances) *base.Instances {
	return nil
}

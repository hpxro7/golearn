package cluster

import (
	"fmt"
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

func (kMeans *KMeans) Fit(data *base.Instances) {
	if data.Rows < kMeans.Clusters {
		panic("kmeans: Fewer training instances than clusters")
	}
	kMeans.TrainingData = data
	centroids := kMeans.init(kMeans)
	fmt.Println(centroids)
}

func (kMeans *KMeans) Predict(data *base.Instances) *base.Instances {
	return nil
}

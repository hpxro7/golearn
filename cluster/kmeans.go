package cluster

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hpxro7/golearn/base"
)

type kMeansCluster struct {
	k            int
	trials       int
	tolerance    float64
	trainingData *base.Instances
}

// Defines cluster initialization types
type InitType int8

const (
	KMeansPlus InitType = iota
	Randomized
)

func NewKMeansCluster(clusters, initTrials int, tolerance float64) *kMeansCluster {
	cluster := &kMeansCluster{
		k:         clusters,
		trials:    initTrials,
		tolerance: tolerance,
	}
	return cluster
}

func (kmeans *kMeansCluster) Fit(data *base.Instances) {
	kmeans.trainingData = data
}

func (kmeans *kMeansCluster) Predict(data *base.Instances) *base.Instances {
	centroids := kmeans.initCentroids(data, Randomized)
	return nil
}

func (kmeans *kMeansCluster) initCentroids(data *base.Instances, init InitType) map[int][]float64 {
	centroids := make(map[int][]float64, kmeans.k)
	switch init {
	case KMeansPlus:
		panic("K-means++ Initialization not Supported...")
	case Randomized:
		{
			rand.Seed(time.Now().Unix())
			randomRows := rand.Perm(data.Rows)[:kmeans.k]
			for i, row := range randomRows {
				fmt.Println(row)
				centroids[i] = data.GetRowVector(row)
			}
		}
	default:
		panic("Unknown initialization type used")
	}

	return centroids
}

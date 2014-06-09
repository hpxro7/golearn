package main

import (
	"fmt"

	"github.com/hpxro7/golearn/base"
	"github.com/hpxro7/golearn/cluster"
)

func main() {
	data, err := base.ParseCSVToInstances("../datasets/clusterdata.csv", false)
	if err != nil {
		errString := fmt.Sprint("Parsing error has occurred - ", err)
		panic(errString)
	}

	/*chooseFirstK := func(kmeans *cluster.KMeans) [][]float64 {
		data := kmeans.TrainingData
		var centroids [][]float64
		for i := 0; i < kmeans.Clusters; i++ {
			centroids = append(centroids, data.GetRowVector(i))
		}
		return centroids
	}*/

	kCluster := cluster.NewKMeans(3, 5, cluster.KRandomized, 1e-2)
	kCluster.Fit(data)
	kCluster.Predict(data)

	fmt.Println("Number of rows: ", data.Rows)
}

package kmeans

import (
	"go-core"
	"go-core/maths"
	"math"
	"math/rand"
)

/**
Options for the k-menas segmentation algorithm
*/
type KmeansOptions struct {
	/**
	Distance function to use
	Default: SSD (sum squared)
	*/
	Distance func(a maths.Vector, b maths.Vector) float64

	/**
	Range of values for each of the matrix items
	Used to create a random Centroid,
	Default: 0 - 2^64
	*/
	ValuesRange [][2]float64 // todo Implementation

	/**
	Number of Tries to approximate the centroids each time with new random ones
	Default: 3
	*/
	Tries int

	/**
	Maximum iterations
	*/
	MaxIter int
}

func Kmeans(input maths.Matrix, k int, options KmeansOptions) []maths.Matrix {
	options.NormalizeOptions() // normalize options

	t := options.Tries
	var arrCentroids []maths.Matrix

	ch := make(chan maths.Matrix, t)

	for i := 0; i < options.Tries; i++ { //todo - Manage routines
		go Segment(input, k, options, ch)
	}
	for i := 0; i < options.Tries; i++ {
		arrCentroids = append(arrCentroids, <-ch)
	}

	return arrCentroids
}

func Segment(input maths.Matrix, k int, options KmeansOptions, ch chan<- maths.Matrix) {
	centroids := CreateCentroids(k, len(input[0]))
	iter := options.MaxIter
	for core.InRange(&iter) {
		centroidClusters := GetClusters(centroids, input, options.Distance)
		newCentroids := maths.Matrix{}
		for i := range centroidClusters {
			// We remove the Centroid if the Cluster size associated to it is zero (i.e. not associated to a Cluster)
			if len(centroidClusters[i].Cluster) != 0 {
				newCentroids = append(newCentroids, GetNewCentroid(centroidClusters[i].Cluster)) // get new Centroid from Cluster
			}
		}
		// update centroids
		centroids = newCentroids
	}
	ch <- centroids
}

/**
Normalize options passed by user
*/
func (options *KmeansOptions) NormalizeOptions() {
	// default options
	var DEF_OPTS KmeansOptions = KmeansOptions{
		Distance:    SumSquaredDistance,
		ValuesRange: nil,
		Tries:       3,
		MaxIter:     100}

	if (*options).Distance == nil {
		(*options).Distance = DEF_OPTS.Distance
	}

	if (*options).ValuesRange == nil {
		(*options).ValuesRange = DEF_OPTS.ValuesRange
	}

	if (*options).Tries == 0 {
		(*options).Tries = DEF_OPTS.Tries
	}

	if (*options).MaxIter == 0 {
		(*options).MaxIter = DEF_OPTS.MaxIter
	}
}

/**
Gets the Cluster of points linked to associated centroids
*/
func GetClusters(centroids maths.Matrix, points maths.Matrix, distance func(vectorA maths.Vector, vectorB maths.Vector) float64) CentroidClusters {
	ccs, _ := (&CentroidClusters{}).New(centroids, nil) // Create new Centroid Clusters struct

	for p := range points {
		var centroid maths.Vector = centroids[0]
		var minDistance float64 = distance(centroids[0], points[p]) // initialize minimum Distance
		for i := 1; i < len(centroids); i++ {
			d := distance(centroids[i], points[p])
			if d < minDistance {
				minDistance = d
				centroid = centroids[i]
			}
		}
		_, _ = ccs.AssignSingle(centroid, points[p])
	}
	return ccs
}

/**
Generate centroids
k int : number of centroids
n int : size of matrix (len(input[0]))
*/
func CreateCentroids(k int, n int) maths.Matrix {
	var centroids maths.Matrix
	i := k

	// Generate centroids
	for core.InRange(&i) {
		j := n
		centroid := []float64{}
		for core.InRange(&j) {
			centroid[j] = rand.Float64()
		}
		centroids = append(centroids, centroid)
	}
	return centroids
}

/**
Gets the mean of a Cluster to get the
*/
func GetNewCentroid(cluster maths.Matrix) maths.Vector {
	pixels := maths.Transpose(cluster)
	centroid := maths.Vector{}
	for i := range pixels {
		centroid[i] = maths.Mean(pixels[i])
	}
	return centroid
}

/**
Sum Squared Distances
vectors a and b need to be the same size
*/
func SumSquaredDistance(a maths.Vector, b maths.Vector) float64 {
	n := len(a)
	var sum float64
	for core.InRange(&n) {
		sum += maths.Square(math.Abs(float64(b[n] - a[n])))
	}
	return sum
}

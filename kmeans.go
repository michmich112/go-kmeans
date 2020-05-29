package go_kmeans

import (
	"core"
	"core/maths"
	"math"
	"math/rand"
)


/**
Options for the k-menas segmentation algorithm
*/
type kmeansOptions struct {
	/**
	Distance function to use
	Default: SSD (sum squared)
	*/
	distance func(a maths.Matrix, b maths.Matrix) float64

	/**
	Range of values for each of the matrix items
	Used to create a random centroid,
	Default: 0 - 2^64
	*/
	valuesRange [][2]float64

	/**
	Number of tries to approximate the centroids each time with new random ones
	Default: 3
	*/
	tries int

	/**
	Maximum iterations
	*/
	maxIter int
}

/**
Centroid Cluster struct linking a centroid to a cluster of datapoints
 */
type CentroidCluster struct {
	/**
	Centroid point
	 */
	centroid maths.Vector

	/**
	Cluster of points
	 */
	cluster maths.Matrix
}

func kmeans(input maths.Matrix, k int, options kmeansOptions) matrix {

	// default options
	/*var DEF_OPTS kmeansOptions = kmeansOptions{
		distance:    ssd,
		valuesRange: nil,
		tries:       3,
		maxIter:     100}*/

	t := options.tries
	for core.InRange(&t){
		centroids := CreateCentroids(k, len(input[0]))
		iter := options.maxIter
		for core.InRange(&iter) {
			// TODO Implement
		}

	}


	return input
}

// TODO: Implement
func GetClusters(centroids []CentroidCluster) {
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
Gets the mean of a cluster to get the
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
		sum += maths.Square(math.Abs(float64(b[n]-a[n])))
	}
	return sum
}

package kmeans

import (
	"github.com/gildas/go-errors"
	"go-core"
	"go-core/maths"
	"math"
	"math/rand"
	"net"
)

/**
Options for the k-menas segmentation algorithm
*/
type kmeansOptions struct {
	/**
	Distance function to use
	Default: SSD (sum squared)
	*/
	distance func(a maths.Vector, b maths.Vector) float64

	/**
	Range of values for each of the matrix items
	Used to create a random centroid,
	Default: 0 - 2^64
	*/
	valuesRange [][2]float64 // todo Implementation

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

/**
Array of Centroid Cluster structs
*/
type CentroidClusters []CentroidCluster

/**
Created a new array of Centroid Clusters / fills up an existing one with the centroid and cluster data passed in by the user
The `cluster` parameter is initialized to the maths.Matrix zero value by default
Size of the new CentroidClusters is dependent on the size of the centroid Matrix.
Excess cluster data is ignored (i.e. if centroid has 3 elements and data has 4, the last element is ignored and not put in the array)
*/
func (ccs *CentroidClusters) New(centroids maths.Matrix, data []maths.Matrix) (ccss CentroidClusters, err error) {
	// if(centroids == nil && data == nil) -> do nothing
	nccs := CentroidClusters{} // New CentroidClusters
	errText := ""              // Stores the warning/error text

	if centroids != nil {
		for c := range centroids {
			// Add a new Centroid Cluster
			nccs = append(nccs, CentroidCluster{
				centroid: centroids[c],
				cluster:  maths.Matrix{}})
		}
		if data != nil && len(centroids) > len(data) {
			for d := range data {
				nccs[d].cluster = data[d]
			}
			errText += "Less data than centroids.\n"
		} else if data != nil {
			for c := range centroids {
				nccs[c].cluster = data[c]
			}
			errText += "More data than centroids, excess data ignored.\n"
		} else {
			errText += "No cluster data.\n"
		}
	} else {
		errText += "No Centroid Data. Returning empty CentroidClusters.\n"
	}

	*ccs = nccs
	if errText == "" {
		return nccs, nil
	} else {
		return nccs, errors.New(errText)
	}
}

/**
Adds a single vector to the cluster matrix of the assigned centroid
Returns Argument Invalid error if centroid was not found
*/
func (ccs *CentroidClusters) AssignSingle(centroid maths.Vector, data maths.Vector) (ccss CentroidClusters, err error) {
	for i := range *ccs {
		if ((*ccs)[i]).centroid.Equals(centroid) {
			cluster := ((*ccs)[i]).cluster
			if cluster == nil {
				((*ccs)[i]).cluster = maths.Matrix{data}
			} else {
				((*ccs)[i]).cluster = append(cluster, data)
			}
			return *ccs, nil
		}
	}
	return *ccs, errors.ArgumentInvalid.WithMessage("Centroid not found in centroid cluster array.")
}

func kmeans(input maths.Matrix, k int, options kmeansOptions) maths.Matrix {

	// default options
	var DEF_OPTS kmeansOptions = kmeansOptions{
		distance:    SumSquaredDistance,
		valuesRange: nil,
		tries:       3,
		maxIter:     100}

	t := options.tries
	var arrCentroids []maths.Matrix

	for core.InRange(&t) { //todo - seperate into goroutinges
		centroids := CreateCentroids(k, len(input[0]))
		iter := options.maxIter
		for core.InRange(&iter) {
			centroidClusters := GetClusters(centroids, input, options.distance)
			newCentroids := maths.Matrix{}
			for i := range centroidClusters {
				// We remove the centroid if the cluster size associated to it is zero (i.e. not associated to a cluster)
				if len(centroidClusters[i].cluster) != 0 {
					newCentroids = append(newCentroids, GetNewCentroid(centroidClusters[i].cluster)) // get new centroid from cluster
				}
			}
			// update centroids
			centroids = newCentroids
		}
		arrCentroids = append(arrCentroids, centroids)
	}

	return input
}

/**
Gets the cluster of points linked to associated centroids
*/
func GetClusters(centroids maths.Matrix, points maths.Matrix, distance func(vectorA maths.Vector, vectorB maths.Vector) float64) CentroidClusters {
	ccs, _ := (&CentroidClusters{}).New(centroids, nil) // Create new Centroid Clusters struct

	for p := range points {
		var centroid maths.Vector = centroids[0]
		var minDistance float64 = distance(centroids[0], points[p]) // initialize minimum distance
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
		sum += maths.Square(math.Abs(float64(b[n] - a[n])))
	}
	return sum
}

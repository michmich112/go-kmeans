package kmeans

import (
	"github.com/gildas/go-errors"
	"go-core/maths"
)

/* --------------------- CENTROID CLUSTER --------------------- */

/**
Centroid Cluster struct linking a Centroid to a Cluster of datapoints
*/
type CentroidCluster struct {
	/**
	Centroid point
	*/
	Centroid maths.Vector

	/**
	Cluster of points
	*/
	Cluster maths.Matrix
}

/**
Verifies that two Centroid clusters are equal
*/
func (cc *CentroidCluster) Equals(cc2 CentroidCluster) bool {
	return (*cc).Centroid.Equals(cc2.Centroid) && (*cc).Cluster.Equals(cc2.Cluster)
}

/* --------------------- CENTROID CLUSTERS --------------------- */

/**
Array of Centroid Cluster structs
*/
type CentroidClusters []CentroidCluster

/**
Created a new array of Centroid Clusters / fills up an existing one with the Centroid and Cluster data passed in by the user
The `Cluster` parameter is initialized to the maths.Matrix zero value by default
Size of the new CentroidClusters is dependent on the size of the Centroid Matrix.
Excess Cluster data is ignored (i.e. if Centroid has 3 elements and data has 4, the last element is ignored and not put in the array)
*/
func (ccs *CentroidClusters) New(centroids maths.Matrix, data []maths.Matrix) (ccss CentroidClusters, err error) {
	// if(centroids == nil && data == nil) -> do nothing
	nccs := CentroidClusters{} // New CentroidClusters
	errText := ""              // Stores the warning/error text

	if centroids != nil {
		for c := range centroids {
			// Add a new Centroid Cluster
			nccs = append(nccs, CentroidCluster{
				Centroid: centroids[c],
				Cluster:  maths.Matrix{}})
		}
		if data != nil && len(centroids) > len(data) {
			for d := range data {
				nccs[d].Cluster = data[d]
			}
			errText += "Less data than centroids.\n"
		} else if data != nil {
			for c := range centroids {
				nccs[c].Cluster = data[c]
			}
			errText += "More data than centroids, excess data ignored.\n"
		} else {
			errText += "No Cluster data.\n"
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
Delete a Centroid Clusters
*/
func (ccs CentroidClusters) Delete(index int) CentroidClusters {
	length := len(ccs)
	ccs[index] = ccs[length-1]        // Copy last element to index i.
	ccs[length-1] = CentroidCluster{} // Erase last element (write zero value).
	ccs = ccs[:length-1]              // Truncate slice.
	return ccs
}

/**
Adds a single vector to the Cluster matrix of the assigned Centroid
Returns Argument Invalid error if Centroid was not found
*/
func (ccs CentroidClusters) AssignSingle(centroid maths.Vector, data maths.Vector) (ccss CentroidClusters, err error) {
	for i := range ccs {
		if (ccs[i]).Centroid.Equals(centroid) {
			cluster := (ccs[i]).Cluster
			if cluster == nil {
				(ccs[i]).Cluster = maths.Matrix{data}
			} else {
				(ccs[i]).Cluster = append(cluster, data)
			}
			return ccs, nil
		}
	}
	return ccs, errors.ArgumentInvalid.WithMessage("Centroid not found in Centroid Cluster array.")
}

/**
Verifies that two Centroid Clusters are the same (deep Equal)
*/
func (ccs CentroidClusters) Equals(ccss CentroidClusters) bool {
	if len(ccs) != len(ccss) {
		return false
	}

	var indexes []int
	for cc := range ccs {
		var match bool = false
		for i := range ccss {
			if ccs[cc].Equals(ccss[i]) && !contains(indexes, i) {
				match = true
				indexes = append(indexes, i)
			}
		}
		if !match {
			return false
		}
	}
	return true
}

/**
Returns a truth value based if the value is in the slice/array
*/
func contains(arr []int, val int) bool {
	for _, i := range arr {
		if i == val {
			return true
		}
	}
	return false
}

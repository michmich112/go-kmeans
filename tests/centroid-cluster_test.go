package tests

import (
	"go-core/maths"
	"go-kmeans"
	"testing"
)

/* ------------ CENTROID CLUSTER TESTS ------------ */

func TestCentroidClusterEquals(t *testing.T) {
	cc1 := kmeans.CentroidCluster{
		maths.Vector{1, 2, 3},
		maths.Matrix{
			maths.Vector{4, 5, 6},
		},
	}
	cc2 := kmeans.CentroidCluster{
		maths.Vector{1, 2, 3},
		maths.Matrix{
			maths.Vector{4, 5, 6},
		},
	}
	cc3 := kmeans.CentroidCluster{
		maths.Vector{1, 2, 3},
		maths.Matrix{
			maths.Vector{4, 6},
		},
	}
	cc4 := kmeans.CentroidCluster{
		maths.Vector{1, 2},
		maths.Matrix{
			maths.Vector{4, 5, 6},
		},
	}

	// Expect equal
	if !cc1.Equals(cc2) {
		t.Error("Expected to be equal.")
	}

	// Expect not equal
	if cc1.Equals(cc3) {
		t.Error("Expected to be not equal.")
	}

	// Expect not equal
	if cc1.Equals(cc4) {
		t.Error("Expected to be not equal.")
	}

}

/* --------------- CENTROID CLUSTERS --------------- */

func TestCentroidClustersEquals(t *testing.T) {
	ccs1 := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5, 6},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
				maths.Vector{0, 1, 1},
			}},
	}

	ccs2 := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5, 6},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
				maths.Vector{0, 1, 1},
			}},
	}

	ccs3 := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
	}

	ccs4 := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5, 6},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
				maths.Vector{0, 1, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{7, 8, 9},
			Cluster: maths.Matrix{
				maths.Vector{1, 0, 0},
				maths.Vector{1, 0, 1},
			}},
	}

	ccs5 := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5, 6},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
			}},
	}

	ccs6 := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
				maths.Vector{0, 1, 1},
			}},
	}

	// Expect Equals
	if !ccs1.Equals(ccs2) {
		t.Error("Expected Equals")
	}

	// Expect Not Equals
	if ccs1.Equals(ccs3) {
		t.Error("Expected Not Equals")
	}
	if ccs1.Equals(ccs4) {
		t.Error("Expected Not Equals")
	}
	if ccs1.Equals(ccs5) {
		t.Error("Expected Not Equals")
	}
	if ccs1.Equals(ccs6) {
		t.Error("Expected Not Equals")
	}
}

func TestCentroidClustersDelete(t *testing.T) {
	ccs := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5, 6},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
				maths.Vector{0, 1, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{7, 8, 9},
			Cluster: maths.Matrix{
				maths.Vector{1, 0, 0},
				maths.Vector{1, 0, 1},
			}},
	}

	ccss := kmeans.CentroidClusters{
		kmeans.CentroidCluster{
			Centroid: maths.Vector{1, 2, 3},
			Cluster: maths.Matrix{
				maths.Vector{0, 0, 0},
				maths.Vector{0, 0, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{4, 5, 6},
			Cluster: maths.Matrix{
				maths.Vector{0, 1, 0},
				maths.Vector{0, 1, 1},
			}},
		kmeans.CentroidCluster{
			Centroid: maths.Vector{7, 8, 9},
			Cluster: maths.Matrix{
				maths.Vector{1, 0, 0},
				maths.Vector{1, 0, 1},
			}},
	}

	ccs.Delete(1)

	equalityCondition := (ccs[0].Equals(ccss[0]) || ccs[0].Equals(ccss[2])) && (ccs[1].Equals(ccss[0]) || ccs[1].Equals(ccss[2]))

	if !equalityCondition {
		t.Error("Did not delete proper element:\n", ccs, "\n", ccss)
	}
}

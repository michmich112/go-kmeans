package tests

import (
	"go-kmeans"
	"testing"
)

/**
 Tests the InRange method
 */
func TestInRange(t *testing.T) {
	var n int
	k := 3

	for go_kmeans.InRange(&k) {
		n++
	}

	if n != 3 {
		t.Error("Values dont match")
	}

}

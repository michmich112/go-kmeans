package tests

import (
	"go-core"
	"testing"
)

/**
Tests the InRange method
*/
func TestInRange(t *testing.T) {
	var n int
	k := 3

	for core.InRange(&k) {
		n++
	}

	if n != 3 {
		t.Error("Values dont match")
	}

}

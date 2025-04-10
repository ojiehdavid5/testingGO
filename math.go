package main

import "testing"

// Average calculates the average of a slice of float64 numbers.
func Average(xs []float64) float64 {
    if len(xs) == 0 {
        return 0 // or you could return an error
    }
    
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// TestAverage tests the Average function.
func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1,2})



	if v != 1.5 {
	t.Error("Expected 1.5, got ", v)
	}
	}
package main

import "testing"

func BenchmarkLinearRegression(b *testing.B) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	for i := 0; i < b.N; i++ {
		_, _, _ = LinearRegression(x, y)
	}
}

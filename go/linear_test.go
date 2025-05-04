package main

import "testing"

func TestLinearRegression(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	slope, intercept, err := LinearRegression(x, y)
	if err != nil {
		t.Fatalf("Regression failed: %v", err)
	}

	expectedSlope := 0.5
	expectedIntercept := 3.0
	epsilon := 0.01

	if slope < expectedSlope-epsilon || slope > expectedSlope+epsilon {
		t.Errorf("Unexpected slope: got %.2f, expected %.2f", slope, expectedSlope)
	}
	if intercept < expectedIntercept-epsilon || intercept > expectedIntercept+epsilon {
		t.Errorf("Unexpected intercept: got %.2f, expected %.2f", intercept, expectedIntercept)
	}
}

// TestLinearRegressionEmpty checks the behavior of LinearRegression with empty slices.
func TestLinearRegressionEmpty(t *testing.T) {
	x := []float64{}
	y := []float64{}

	_, _, err := LinearRegression(x, y)
	if err == nil {
		t.Fatal("Expected error for empty slices, got nil")
	}
}

func TestLinearRegressionCop(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	slope, intercept, err := LinearRegressionCop(x, y)
	if err != nil {
		t.Fatalf("Regression failed: %v", err)
	}

	expectedSlope := 0.5
	expectedIntercept := 3.0
	epsilon := 0.01

	if slope < expectedSlope-epsilon || slope > expectedSlope+epsilon {
		t.Errorf("Unexpected slope: got %.2f, expected %.2f", slope, expectedSlope)
	}
	if intercept < expectedIntercept-epsilon || intercept > expectedIntercept+epsilon {
		t.Errorf("Unexpected intercept: got %.2f, expected %.2f", intercept, expectedIntercept)
	}
}

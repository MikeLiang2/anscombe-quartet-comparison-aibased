package main

import (
	"errors"

	"github.com/montanaflynn/stats"
)

// LinearRegression computes the slope and intercept of the regression line
// for the given x and y values using the least squares method.
// It returns the slope and intercept, or an error if the input is invalid.
func LinearRegression(xVals, yVals []float64) (float64, float64, error) {
	if len(xVals) != len(yVals) {
		return 0, 0, errors.New("x and y must be of same length")
	}

	xMean, _ := stats.Mean(xVals)
	yMean, _ := stats.Mean(yVals)

	var num, den float64
	for i := range xVals {
		num += (xVals[i] - xMean) * (yVals[i] - yMean)
		den += (xVals[i] - xMean) * (xVals[i] - xMean)
	}
	if den == 0 {
		return 0, 0, errors.New("division by zero in regression")
	}
	slope := num / den
	intercept := yMean - slope*xMean
	return slope, intercept, nil
}

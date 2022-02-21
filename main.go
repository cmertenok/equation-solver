package main

import (
	//"fmt"
	"math"
)

func equationSolver(values []float64) []float64 {
	var roots []float64

	var (
		numA = values[0]
		numB = values[1]
		numC = values[2]
	)

	var discriminant = (numB * numB) - (4 * numA * numC)

	if discriminant == 0 {
		var root = -numB / (2 * numA)
		roots = append(roots, root)
	} else if discriminant > 0 {
		var firstRoot = (-numB + math.Sqrt(discriminant)) / (2 * numA)
		var secondRoot = (-numB - math.Sqrt(discriminant)) / (2 * numA)
		roots = append(roots, firstRoot, secondRoot)
	}

	return roots
}

func main() {

}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func interactiveMode() []float64 {

	var param []float64
	var args = []string{"a", "b", "c"}

	var reader = bufio.NewReader(os.Stdin)

	for i := 0; i != 3; i++ {
		fmt.Printf("%s: ", args[i])
		var parameters, err = reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error. Failed to read input. Please try again", err)
			continue
		}

		parameters = strings.TrimSuffix(parameters, "\r\n")
		parameterFloat, err := strconv.ParseFloat(parameters, 64)

		if err != nil {
			fmt.Printf("Error. Expected a valid real number, got %s instead\n", parameters)
			continue
		}

		if parameterFloat == 0 && len(param) == 0 {
			fmt.Print("Error. a cannot be 0\n")
			continue
		}

		param = append(param, parameterFloat)
	}

	return param
}

func main() {

}

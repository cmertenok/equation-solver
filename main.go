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
	var args = []string{"A", "B", "C"}

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

func nonInteractiveMod(fileName string) []float64 {

	var parameters []float64

	file, err := os.ReadFile(fileName) 

	if err != nil {
	  fmt.Println("Error occured while trying to open file", err)
	  os.Exit(1)
	}

	str :=strings.Fields(string(file))

	if len(str) != 3 {
	  fmt.Fprint(os.Stderr, "Error. Wrong number of parameters. Must be 3 \n")
	  os.Exit(1)
	}

	for index, values := range str {
	 params, err := strconv.ParseFloat(values, 64)

	 if err != nil {
	  fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	  os.Exit(1)
	 }

	 if index == 0 && params == 0 {
	  fmt.Print("Error. a cannot be 0\n")
		os.Exit(1)
	 }

	  parameters = append(parameters,params)
	}

	return parameters
 }

func main() {

  var fileName string
  var input []float64 
  var results []float64

  if len(os.Args) > 1 {
    fileName = os.Args[1]
  }

  if fileName != "" {
    input = nonInteractiveMod(fileName)
  } else {
    input = interactiveMode()
  }

  results = equationSolver(input)

  fmt.Printf("Equation: %.2fX^2 + %.2fX + %.2f = 0\n", input[0], input[1], input[2])

  if len(results) == 1 {
    fmt.Printf("X1 = %f\n", results[0])
  } else if len(results) == 2 {
    fmt.Printf("X1 = %f\nX2 = %f\n", results[0], results[1])
  }

  fmt.Printf("The equation has %d roots", len(results))

}

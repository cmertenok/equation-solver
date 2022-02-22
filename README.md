# Equation-solver
This simple CLI application can be used to solve quadratic equations

```
A: 2
B: 1
C: -3
Equation is: (2) x^2 + (1) x + (-3) = 0 
There are 2 roots
x1 = 1.00
x2 = -1.50
```

## Installation

Using this package requires a working Go environment. 

Install using the following command:

```
$ git clone https://github.com/cmertenok/equation-solver.git
```

Make sure to have this repository cloned and open and run `go install`

#### Interactive Mode
To start in interactive mode simply type *go run* command with the name of the source file.

```
go run main.go
```
#### Non-interactive Mode
In non-interactive mode you need to have a file in .txt format containing three numbers (more or less numbers result in an error) divided by spaces. 
Example of file structure:
```
2 1 -3
```
To start the application type the same command as in interactive mode, but add a file name to the end. 
For example test.txt:
```
go run main.go test.txt
```
## Revert commit

[299294c753d64ca78e7242164eb3f2e65817104b](https://github.com/cmertenok/equation-solver/commit/299294c753d64ca78e7242164eb3f2e65817104b) 
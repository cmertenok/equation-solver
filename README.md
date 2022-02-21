# SDM-Quadratic-Equations
Software Development Methods."Console application for solving quadratic equations".

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

## Usage

This application includes two modes: interactive and file modes

To use interactive mode, simply call `sample-app` command and import the parameters of your equation

```
$ equation-solver
a = 2
b = 4
c = 2
```

To use file mode, make sure you have a file containing real numbers with a space between each, like this:

Example:
```
1 2 3
```

Then run the following command:

```
$ sample-appr test.txt
```
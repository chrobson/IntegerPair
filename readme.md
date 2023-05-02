# Pair with Given Sum in Go

This program is an implementation of the problem of finding a pair with a given sum in an unsorted integer array, and presenting the pair with the higher number first.

## How to Use

To use this program, simply run the `main.go` file using a Go compiler, such as `go run main.go`.

Optionally can be run from Go playground : https://go.dev/play/p/ch3bAP_iguX 

The program currently contains an example array and sum in the `main` function for testing purposes. 

The `findPair` function takes in two arguments:

1. `arr []int`: an unsorted integer array
2. `sum int`: an integer sum to find pairs for

The function returns a two-dimensional slice of pairs that add up to the given sum, sorted as described to promt first higher number.


The program does not use the `sort` library to perform the sorting. Instead, it implements a custom sorting algorithm using a nested loop.
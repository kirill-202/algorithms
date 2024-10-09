package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func MultiplyMatrices(matrixA, matrixB [][]int, numWorkers int) [][]int {

	m := len(matrixA)
	p := len(matrixB[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, p)
	}

	ch := make(chan int, numWorkers) 

	wg.Add(m * p)

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			go calculateRowColumn(matrixA[i], extractColumn(matrixB, j), i, j, result, ch)
		}
	}


	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

	return result
}

func calculateRowColumn(row, column []int, i, j int, result [][]int, ch chan int) {
	defer wg.Done()
	sum := 0
	for k := 0; k < len(row); k++ {
		sum += row[k] * column[k]
	}
	result[i][j] = sum
	ch <- sum
}

// Extract a column from matrix B
func extractColumn(matrix [][]int, col int) []int {
	column := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		column[i] = matrix[i][col]
	}
	return column
}

func main() {
	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}
	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}

	fmt.Println("Matrix A:", matrixA)
	fmt.Println("Matrix B:", matrixB)

	result := MultiplyMatrices(matrixA, matrixB, 4)

	fmt.Println("Result matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
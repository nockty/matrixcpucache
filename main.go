package main

import (
	"runtime"
	"sync"
)

func main() {

}

type LinkedMatrixNode struct {
	value   int64
	nextRow *LinkedMatrixNode
	nextCol *LinkedMatrixNode
}

func (m *LinkedMatrixNode) sum() int64 {
	var sum int64
	for row := m; row != nil; row = row.nextRow {
		for col := row; col != nil; col = col.nextCol {
			sum += col.value
		}
	}
	return sum
}

type Matrix struct {
	matrix  [][]int64
	numRows int
	numCols int
}

func (m *Matrix) sumCacheMiss() int64 {
	var sum int64
	for j := 0; j < m.numCols; j++ {
		for i := 0; i < m.numRows; i++ {
			sum += m.matrix[i][j]
		}
	}
	return sum
}

func (m *Matrix) sumCacheHit() int64 {
	var sum int64
	for i := 0; i < m.numRows; i++ {
		for j := 0; j < m.numCols; j++ {
			sum += m.matrix[i][j]
		}
	}
	return sum
}

func (m *Matrix) sumFalseSharing() int64 {
	numThreads := runtime.NumCPU()
	counts := make([]int64, numThreads)
	var wg sync.WaitGroup
	wg.Add(numThreads)
	for k := 0; k < numThreads; k++ {
		go func(k int) {
			for i := k * m.numRows / numThreads; i < (k+1)*m.numRows/numThreads; i++ {
				for j := 0; j < m.numCols; j++ {
					counts[k] += m.matrix[i][j]
				}
			}
			wg.Done()
		}(k)
	}
	wg.Wait()
	var sum int64
	for i := 0; i < numThreads; i++ {
		sum += counts[i]
	}
	return sum
}

func (m *Matrix) sumParallel() int64 {
	numThreads := runtime.NumCPU()
	counts := make(chan int64, numThreads)
	var wg sync.WaitGroup
	wg.Add(numThreads)
	for k := 0; k < numThreads; k++ {
		go func(k int) {
			var count int64
			for i := k * m.numRows / numThreads; i < (k+1)*m.numRows/numThreads; i++ {
				for j := 0; j < m.numCols; j++ {
					count += m.matrix[i][j]
				}
			}
			counts <- count
			wg.Done()
		}(k)
	}
	go func() {
		wg.Wait()
		close(counts)
	}()
	var sum int64
	for count := range counts {
		sum += count
	}
	return sum
}

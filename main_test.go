package main

import "testing"

func Test_sum(t *testing.T) {
	testCases := []struct {
		matrix   [][]int64
		expected int64
	}{
		{
			matrix: [][]int64{
				{1},
			},
			expected: 1,
		},
		{
			matrix: [][]int64{
				{1, 2},
			},
			expected: 3,
		},
		{
			matrix: [][]int64{
				{1},
				{2},
				{3},
			},
			expected: 6,
		},
		{
			matrix: [][]int64{
				{1, 2},
				{3, 4},
			},
			expected: 10,
		},
		{
			matrix: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 45,
		},
		{
			matrix: [][]int64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: 136,
		},
	}
	for _, tc := range testCases {
		linkedMatrix := newLinkedTestMatrix(tc.matrix)
		actual := linkedMatrix.sum()
		if actual != tc.expected {
			t.Errorf("got %d, want %d", actual, tc.expected)
		}

		matrix := newTestMatrix(tc.matrix)
		actual = matrix.sumCacheMiss()
		if actual != tc.expected {
			t.Errorf("got %d, want %d", actual, tc.expected)
		}
		actual = matrix.sumCacheHit()
		if actual != tc.expected {
			t.Errorf("got %d, want %d", actual, tc.expected)
		}
		actual = matrix.sumFalseSharing()
		if actual != tc.expected {
			t.Errorf("got %d, want %d", actual, tc.expected)
		}
		actual = matrix.sumParallel()
		if actual != tc.expected {
			t.Errorf("got %d, want %d", actual, tc.expected)
		}
	}
}

func newTestMatrix(matrix [][]int64) *Matrix {
	return &Matrix{
		matrix:  matrix,
		numRows: len(matrix),
		numCols: len(matrix[0]),
	}
}

func newLinkedTestMatrix(matrix [][]int64) *LinkedMatrixNode {
	numRows := len(matrix)
	numCols := len(matrix[0])
	m := make([][]*LinkedMatrixNode, numRows)
	for i := 0; i < numRows; i++ {
		m[i] = make([]*LinkedMatrixNode, numCols)
		for j := 0; j < numCols; j++ {
			m[i][j] = &LinkedMatrixNode{
				value: matrix[i][j],
			}
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols-1; j++ {
			m[i][j].nextCol = m[i][j+1]
		}
	}
	for i := 0; i < numRows-1; i++ {
		for j := 0; j < numCols; j++ {
			m[i][j].nextRow = m[i+1][j]
		}
	}
	return m[0][0]
}

const SIZE = 512

func setupMatrix() *Matrix {
	m := make([][]int64, SIZE)
	for i := 0; i < SIZE; i++ {
		m[i] = make([]int64, SIZE)
		for j := 0; j < SIZE; j++ {
			m[i][j] = int64(i + j)
		}
	}
	matrix := newTestMatrix(m)
	return matrix
}

func setupLinkedMatrix() *LinkedMatrixNode {
	m := make([][]int64, SIZE)
	for i := 0; i < SIZE; i++ {
		m[i] = make([]int64, SIZE)
		for j := 0; j < SIZE; j++ {
			m[i][j] = int64(i + j)
		}
	}
	matrix := newLinkedTestMatrix(m)
	return matrix
}

var result int64

func Benchmark_sumLinked(b *testing.B) {
	matrix := setupLinkedMatrix()
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = matrix.sum()
	}
	result = r
}

func Benchmark_sumCacheMiss(b *testing.B) {
	matrix := setupMatrix()
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = matrix.sumCacheMiss()
	}
	result = r
}

func Benchmark_sumCacheHit(b *testing.B) {
	matrix := setupMatrix()
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = matrix.sumCacheHit()
	}
	result = r
}

func Benchmark_sumFalseSharing(b *testing.B) {
	matrix := setupMatrix()
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = matrix.sumFalseSharing()
	}
	result = r
}

func Benchmark_sumParallel(b *testing.B) {
	matrix := setupMatrix()
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = matrix.sumParallel()
	}
	result = r
}

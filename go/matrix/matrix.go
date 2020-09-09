// Package matrix implements type and utility routines to
// get rows and coloumns of a matrix
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix to hold a representation of matrix
type Matrix [][]int

// New builds a matrix as an array of array of integers
func New(input string) (Matrix, error) {
	rows := strings.Split(input, "\n")

	m := make(Matrix, len(rows))
	for i, r := range rows {
		vals := strings.Fields(r)
		if i > 0 && len(vals) != len(m[i-1]) {
			return nil, errors.New("uneven rows")
		}

		m[i] = make([]int, len(vals))
		for j, val := range vals {
			v, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			m[i][j] = v
		}
	}

	return m, nil
}

// Rows returns number of rows in matrix
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		rows[i] = make([]int, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}
	return rows
}

// Cols return number of columns in matrix
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := 0; i < len(cols); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(cols[i]); j++ {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

// Set given row, col and value sets the value in cell
func (m Matrix) Set(r, c, val int) (ok bool) {
	if ok = r >= 0 && c >= 0 && r < len(m) && c < len(m[0]); ok {
		m[r][c] = val
	}
	return
}

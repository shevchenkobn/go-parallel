package matrices

import (
	"errors"
	"math/rand"
)

func GetRandomMatr(rows, columns int) [][]int {
	if rows <= 0 || columns <= 0 {
		panic(errors.New("invalid dimensions"))
	}

	matr := make([][]int, rows, rows)
	for i := 0; i < rows; i++ {
		matr[i] = make([]int, columns, columns)
		for j := 0; j < columns; j++ {
			matr[i][j] = rand.Int()
		}
	}
	return matr
}

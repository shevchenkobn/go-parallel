package matrices

import (
	"errors"
	"log"
	"os"
	"runtime"
	"sync"
)

var maxProcs = runtime.GOMAXPROCS(0)

func MultiplyMatrices(first, other [][]int) [][]int {
	if first == nil || other == nil || len(first) == 0 || len(other) == 0 {
		panic(errors.New("either matrix is nil or empty"))
	}
	if len(first[0]) != len(other) {
		panic(errors.New("matrices are incompatible"))
	}
	//logger.Println("start")

	result := make([][]int, len(first), len(first))
	for i := range first {
		result[i] = make([]int, len(other[0]), len(other[0]))
		for j := range other[0] {
			sum := 0
			for k := 0; k < len(other); k++ {
				sum += first[i][k] * other[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func MultiplyMatricesQuick(first, other [][]int) [][]int {
	if first == nil || other == nil || len(first) == 0 || len(other) == 0 {
		panic(errors.New("either matrix is nil or empty"))
	}
	if len(first[0]) != len(other) {
		panic(errors.New("matrices are incompatible"))
	}
	//logger.Println("start")

	result := make([][]int, len(first), len(first))
	for i := range first {
		result[i] = make([]int, len(other[0]), len(other[0]))
		for k := 0; k < len(other); k++ {
			//result[i][j] := 0
			for j := range other[0] {
				result[i][j] += first[i][k] * other[k][j]
			}
			//result[i][j] = sum
		}
	}
	return result
}

var logger = log.New(os.Stdout, "", 0)
func GoMultiplyMatrices(first, other [][]int) [][]int {
	if first == nil || other == nil || len(first) == 0 || len(other) == 0 {
		panic(errors.New("either matrix is nil or empty"))
	}
	if len(first[0]) != len(other) {
		panic(errors.New("matrices are incompatible"))
	}
	//logger.Println("start")

	size := len(first)
	result := make([][]int, size, size)

	var wg sync.WaitGroup
	iLimit := size / maxProcs
	step := func (start, count int, updateWg bool) {
		if updateWg {
			defer wg.Done()
		}
		//logger.Printf("start at %v", start)
		//logger.Printf("%v %v\n", i, i + count)
		for i := start; i < start + count; i++ {
			//logger.Printf("%v %v\n", i, i + count)
			result[i] = make([]int, len(other[0]), len(other[0]))
			for j := range other[0] {
				sum := 0
				for k := 0; k < len(other); k++ {
					sum += first[i][k] * other[k][j]
				}
				result[i][j] = sum
			}
		}
	}

	longestOffset := iLimit * (maxProcs - 1)
	//logger.Printf("longest: %v %v\n", longestOffset, size - longestOffset)
	wg.Add(1)
	go step(longestOffset, size - longestOffset, true)
	for goN := 1; goN < maxProcs - 1; goN++ {
		wg.Add(1)
		go step(goN * iLimit, iLimit, true)
	}
	step(0, iLimit, false)

	wg.Wait()
	return result
}

func GoMultiplyMatricesQuick(first, other [][]int) [][]int {
	if first == nil || other == nil || len(first) == 0 || len(other) == 0 {
		panic(errors.New("either matrix is nil or empty"))
	}
	if len(first[0]) != len(other) {
		panic(errors.New("matrices are incompatible"))
	}
	//logger.Println("start")

	size := len(first)
	result := make([][]int, size, size)

	var wg sync.WaitGroup
	iLimit := size / maxProcs
	step := func (start, count int, updateWg bool) {
		if updateWg {
			defer wg.Done()
		}
		//logger.Printf("start at %v", start)
		//logger.Printf("%v %v\n", i, i + count)
		for i := start; i < start + count; i++ {
			//logger.Printf("%v %v\n", i, i + count)
			result[i] = make([]int, len(other[0]), len(other[0]))
			for k := 0; k < len(other); k++ {
				//sum := 0
				for j := range other[0] {
					result[i][j] += first[i][k] * other[k][j]
				}
				//result[i][j] = sum
			}
		}
	}

	longestOffset := iLimit * (maxProcs - 1)
	//logger.Printf("longest: %v %v\n", longestOffset, size - longestOffset)
	wg.Add(1)
	go step(longestOffset, size - longestOffset, true)
	for goN := 1; goN < maxProcs - 1; goN++ {
		wg.Add(1)
		go step(goN * iLimit, iLimit, true)
	}
	step(0, iLimit, false)

	wg.Wait()
	return result
}
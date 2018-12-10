package main

import (
	"./matrices"
	"fmt"
	"reflect"
	"time"
)

const size = 512

func main()  {
	fmt.Printf("Generating matrices of size %v...\n", size)

	first := matrices.GetRandomMatr(size, size)
	fmt.Println("First ready.")
	other := matrices.GetRandomMatr(size, size)
	fmt.Println("Second ready. Multiplying...")
	fmt.Println()

	t0 := time.Now()
	resSingle := matrices.MultiplyMatrices(first, other)
	fmt.Printf("Single-threaded time spent: %v\n", time.Now().Sub(t0))

	t0 = time.Now()
	resMulti := matrices.GoMultiplyMatrices(first, other)
	fmt.Printf("Go-routines time spent: %v\n", time.Now().Sub(t0))
	fmt.Printf("Results are equal: %v\n\n", reflect.DeepEqual(resSingle, resMulti))
}
package main

import (
	"fmt"
	"time"
	"reflect"
	"./matrices"
)

const size = 1027

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
	resQuick := matrices.MultiplyMatricesQuick(first, other)
	fmt.Printf("Single-threaded Quick time spent: %v\n", time.Now().Sub(t0))

	t0 = time.Now()
	resMulti := matrices.GoMultiplyMatrices(first, other)
	fmt.Printf("Go-routines time spent: %v\n", time.Now().Sub(t0))

	t0 = time.Now()
	resMultiQuick := matrices.GoMultiplyMatricesQuick(first, other)
	fmt.Printf("Go-routines Quick time spent: %v\n", time.Now().Sub(t0))

	fmt.Printf(
		"Results are equal: %v, %v, %v\n\n",
		reflect.DeepEqual(resSingle, resMulti),
		reflect.DeepEqual(resSingle, resQuick),
		reflect.DeepEqual(resMulti, resMultiQuick),
	)
}


package main
import (
	"fmt"
	"errors"
)

func MultiplyMatrices(first, other [][]int) {
	if first == nil || other == nil || len(first) == 0 || len(other) == 0 {
		panic(errors.New("either matrix is nil or empty"))
	}
	if len(first[0]) != len(other) {
		panic(errors.New("matrices are incompatible"))
	}

	for i, firstRow := range first {
		otherRow := other[i]
		for j := range other {
			sum := 0

		}
	}
}

func addValue(s *[]int) {
	d := append(*s, 3)
	(*s)[0] = 0
	fmt.Printf("In addValue: s is %v\n", &s)
	fmt.Printf("%v\n", s==&d)
}

func main() {
	s := make([]int, 2, 3)
	fmt.Printf("In main, before addValue: s is %v %v\n", s, len(s))
	addValue(&s)
	fmt.Printf("In main, after addValue: s is %v\n", s)
}

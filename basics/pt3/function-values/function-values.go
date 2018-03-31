package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypoth := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
    fmt.Println(hypoth(5, 12))

	fmt.Println(compute(hypoth))
	fmt.Println(compute(math.Pow))
}

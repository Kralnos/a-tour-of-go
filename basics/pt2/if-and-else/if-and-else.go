package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
    // Both calls to pow are executed and return before the call to fmt.Println in main begins.
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

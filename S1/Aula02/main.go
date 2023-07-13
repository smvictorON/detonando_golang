package main

import (
	"fmt"
	"math"
)

func delta(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func calculate(a float64, b float64, c float64) (bool, float64, float64) {
	delta := delta(a, b, c)
	if delta < 0 {
		return false, math.NaN(), math.NaN()
	} else if delta == 0 {
		x := -b / (2.0 * a)
		return true, x, x
	} else {
		x1 := (-b + math.Sqrt(delta)) / 2 * a
		x2 := (-b - math.Sqrt(delta)) / 2 * a
		return true, x1, x2
	}
}

func main() {
	//baskhara
	roots, x1, x2 := calculate(7, 3, 4)
	fmt.Printf("roots: %t, x1: %v, x2: %v \n", roots, x1, x2)

	roots, x1, x2 = calculate(4, -4, 1)
	fmt.Printf("roots: %t, x1: %v, x2: %v \n", roots, x1, x2)

	roots, x1, x2 = calculate(1, -5, 6)
	fmt.Printf("roots: %t, x1: %v, x2: %v \n", roots, x1, x2)

}

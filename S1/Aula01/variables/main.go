package main

import (
	"fmt"
)

//global
var v1 = 1

func main() {
	//local
	var v2 = 2

	fmt.Println(v1)
	fmt.Println(v2)

	//loop and array
	notas := [5]int{2, 4, 6, 8, 10}
	sumValue := 0
	sumIndex := 0
	for index, nota := range notas {
		sumIndex += index
		sumValue += nota
	}

	fmt.Println(sumIndex)
	fmt.Println(sumValue)
}

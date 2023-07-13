package main

import (
	"fmt"
)

// o problema proposto é, dado um vetor, e um numero informado, quero que me
// retorno os indices dos numeros do vetor que somados seja igual ao numero
// informado sem repetir o mesmo numero no vetor duas vezes

type Solution struct{}

func (s Solution) twoSum(nums []int, target int) []int {
	//transform o array de int em um "dicionario"
	dictValues := make(map[int]int)
	for index, val := range nums {
		dictValues[val] = index
	}
	fmt.Println(dictValues)

	for index, val := range nums {
		complement := target - val
		//no caso 9 - 2 é 7 se existe o dictValues[7] e o index dele é diferente do que
		//eu estou então essa é a resposta
		if _, ok := dictValues[complement]; ok && dictValues[complement] != index {
			return []int{index, dictValues[complement]}
		}
	}

	return nil
}

func main() {
	var sol Solution
	var nums = []int{2, 7, 11, 15}
	var result []int = sol.twoSum(nums, 9)
	fmt.Println("Resultado: ", result)
}

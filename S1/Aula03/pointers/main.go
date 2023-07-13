package main

import (
	"fmt"
)

// O operador * é usado para declarar um tipo de ponteiro ou para acessar o valor apontado por um ponteiro existente
// O Operador & é usado para obter o endereço de memória da variável

func main() {
	var num int = 32
	//aqui estamos criando uma var do tipo "ponteiro de int" que aponta para um endereço de memoria int
	var ptr *int = &num

	fmt.Println(ptr)
	//para acessar o valor para onde esse ponteiro de int está apontando basta colocar *
	fmt.Println(*ptr)
}

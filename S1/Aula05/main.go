package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//faço um reader para ler o que vai ser digitado
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Name:")
	//coloco em text o que foi digitado
	text, _ := reader.ReadString('\n')
	//print o que foi digitado incluindo o enter
	fmt.Printf("Hello, %v how are you today?\n", text)
	//recebo em newName o que foi digitado sem os espaço
	newName := fmt.Sprintf("%v", strings.TrimSpace(text))
	//printo sem os espaços e o enter
	fmt.Printf("Now, %v, without the extra line\n", newName)

	fmt.Println("----")

	fmt.Print("Now, type an integer:")
	nints, _ := reader.ReadString('\n')
	nint, errInt := strconv.ParseInt(strings.TrimSpace(nints), 10, 64)
	if errInt != nil {
		log.Fatal("Invalid Integer")
	}
	fmt.Printf("You typed: %d\n", nint)

	fmt.Println("----")

	fmt.Print("Now, type an float:")
	nfs, _ := reader.ReadString('\n')
	nf, errF := strconv.ParseFloat(strings.TrimSpace(nfs), 64)
	if errF != nil {
		log.Fatal("Invalid float")
	}
	fmt.Printf("You typed: %f\n", nf)

	fmt.Println("----")

	//coverte um float 1.2 em um numero com virgula 1,2
	p := message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("%f", nf)

	//criando um arquivo
	stringsArr := []byte("Good morning.\nEat an Apple!\nHave a nice day!\n")
	//permissão -rw-r--r--
	err := ioutil.WriteFile("./S1/AULA05/example.txt", stringsArr, 0644)
	check(err)

	//lendo o arquivo
	data, err := ioutil.ReadFile("./S1/AULA05/example.txt")
	check(err)
	//T maiusculo imprime o tipo do dado
	fmt.Printf("\nType: %T\n", data)
	//passando o que foi lido para string
	textContent := string(data)
	fmt.Println(textContent)
}

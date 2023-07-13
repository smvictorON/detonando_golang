package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func fibonacci(n int, aChannel chan string) {
	if n <= 1 {
		aChannel <- fmt.Sprintf("The term %d is %d\n", n, n)
		return
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	aChannel <- fmt.Sprintf("The term %d is %d", n, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	chAnswer := make(chan string)
	resultReceived := false

	fmt.Print("Type a term number or other character to finish: ")
	nints, _ := reader.ReadString('\n')
	nint, err := strconv.ParseInt(strings.TrimSpace(nints), 10, 64)
	if err != nil {
		log.Fatal("Invalid Integer")
	}

	// cria um ticker de 1 segundo
	ticker := time.Tick(1 * time.Second)

	start := time.Now()
	go fibonacci(int(nint), chAnswer)

	for !resultReceived {
		select {
		case answer := <-chAnswer:
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Println("Got an answer:", answer)
			fmt.Printf("Time to get the answer: %s\n", elapsed)
			// define a variável de controle para true e sai do loop
			resultReceived = true
		case <-ticker:
			fmt.Println("Waiting...")
		}
	}
}

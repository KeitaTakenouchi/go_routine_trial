package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		exec(i)
	}
}

func exec(n int) {
	round := 10

	a := make(chan string, 10)
	b := make(chan string, 10)
	c := make(chan string, 10)

	for i := 0; i < round; i++ {
		go run(a, "AAA")
		go run(b, "BBB")
		go run(c, "CCC")
	}

	for i := 0; i < 3*round; i++ {
		select {
		case <-a:
			fmt.Print("A")
		case <-b:
			fmt.Print("B")
		case <-c:
			fmt.Print("C")
		}
	}
	fmt.Println()
}

func run(c chan<- string, str string) {
	c <- str
}

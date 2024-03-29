package main

import (
	"fmt"
	"time"
	"math/rand"
)


func boring(msg string) <- chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <- chan string) <- chan string {
	c := make(chan string)
	go func() { 
		for {
			select {
			case s := <- input1: c <- s
			case s := <- input2: c <- s
			}
		}

	 }()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<- c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

//https://www.youtube.com/watch?v=f6kdp27TYZs
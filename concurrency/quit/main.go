package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string, quit chan string) <- chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c)}
	quit <- "Bye!"
	fmt.Println("Joe says: %q\n", <-quit)

	for {
		select {
		case c <- fmt.Sprintf("%s: %d", msg, i):
			// nothing
		case <-quit: 
			// cleanup()
			quit <- "See you!"
			return
		}
	}
}
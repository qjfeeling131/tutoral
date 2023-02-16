package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c := make(chan string)
// 	go boring("boring!", c)
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("You say : %q\n", <-c)
// 	}
// 	fmt.Println("You're boring: I'm leaving")
// }

// func boring(msg string, c chan string) {
// 	for i := 0; ; i++ {
// 		c <- fmt.Sprintf("%s %d", msg, i)
// 		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 	}
// }

type Ball struct {
	hits int
}

func main() {
	// table := make(chan *Ball)
	// go player("ping", table)
	// go player("pong", table)
	// println("Getting start new ball")
	// newBall := new(Ball)
	// fmt.Println("new ball addr: %s", &newBall)
	// table <- newBall

	// time.Sleep(1 * time.Second)
	// <-table
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "ch2111"
		ch2 <- "ch1222"
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		fmt.Println((&ball))
		fmt.Println("this is a new ball? ", ball.hits)
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	// go f()
	// c <- 0
	// fmt.Print(c)
	// fmt.Println(a)
	// for _, w := range work {
	// 	go func(w func()) {
	// 		limit <- 1
	// 		w()
	// 		<-limit
	// 	}(w)
	// }
	// select {}
	// fmt.Println(a)

	//Set properites of the predefined logger,
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"pari", "Gladys", "Jack"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// message, err := greetings.Hello("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(messages)
}

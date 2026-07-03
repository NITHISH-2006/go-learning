package main

import (
	"fmt"
	"time"
)

func player(name string, court chan string) {
	ball := <-court

	fmt.Printf("%s hit the ball!\n", name)

	time.Sleep(500 * time.Millisecond)

	court <- ball
	player(name, court)
}

func main() {

	table := make(chan string)

	go player("Ping", table)
	go player("Pong", table)

	fmt.Println("Game on! Serve the ball")

	table <- "ball"

	select {}

}

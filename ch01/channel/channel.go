package main

import "fmt"

func main() {
	messages := make(chan string, 10)
	messages <- "Leader"
	messages <- "Follower"
	fmt.Println(<-messages, <-messages)
}

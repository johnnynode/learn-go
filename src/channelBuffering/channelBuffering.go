package main

import "fmt"

func main() {
	messages := make(chan stirng, 2)

	messages <- "buffered"

	messages <- "channelSynchronization"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

/*


$ go run channelBuffering.go
buffered
channelSynchronization

*/
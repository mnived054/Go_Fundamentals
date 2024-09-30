package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	sig := <-sigChan
	fmt.Println("\n Received signal from OS:", sig)
}

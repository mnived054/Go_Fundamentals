package main

import "fmt"

func main() {
	var mychan chan int
	fmt.Println("Value of channel: ", mychan)
	fmt.Printf("Type of channel: %T ", mychan)
	mychannel1 := make(chan int)
	fmt.Println("\nValue of Channel1: ", mychannel1)
	fmt.Printf("Type of mychannel1: %T", mychannel1)

}

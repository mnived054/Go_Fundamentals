package main

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- 42
// 	}()
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- "Hi, Golang!"
// 	}()

// 	select {
// 	case value := <-ch1:
// 		fmt.Println("Received from ch1: ", value)
// 	case data := <-ch2:
// 		fmt.Println("Received from ch2:", data)
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
//     c := make(chan string)
//     go boring("boring!", c)
//     for i := 0; i < 5; i++ {
//         fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
//     }
//     fmt.Println("You're boring; I'm leaving.")
// }

// func boring(msg string, c chan string) {
//     for i := 0; ; i++ {
//         c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
//         //time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
//     }
// }

// func main() {
// 	c := boring("boring!") // Function returning a channel.
// 	for i := 0; i < 6; i++ {
// 		fmt.Printf("You say: %q\n", <-c)
// 	}
// 	fmt.Println("You're boring; I'm leaving.")
// }

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	fmt.Println("You're boring; I'm leaving.1111")
	return c // Return the channel to the caller.
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

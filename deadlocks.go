package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := range c1 {
			println("G1 got this from C1", i)
			c2 <- i
		}
	}()

	go func() {
		for i := range c2 {
			println("G2 got", i)
			c1 <- i
		}
	}()

	c1 <- 1
	time.Sleep(1 * 52)
	c1 <- 2
	time.Sleep(1 * 10000)

}

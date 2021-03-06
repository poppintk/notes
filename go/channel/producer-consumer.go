package main

import "time"
import "math/rand"
// When is channel is useful?
// producer and consumer case 

func main() {
	c := make(chan int)
	
	// producer
	for i := 0; i < 4; i++ {
		go doWork(c)
	}
	
	// consumer
	for {
		v := <-c
		println(v)
	}
	// NOTE: we can also achive this by using waitGroup, wait for all threads are finished
}

func doWork(c chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		c <- rand.Int()
	}
}

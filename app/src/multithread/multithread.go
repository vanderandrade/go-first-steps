package multithread

import (
	"fmt"
	"time"
)

func OldFor() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 3)
	}
}

func ParallelFor() {
	channel := make(chan int)

	for i := 1; i <= 5; i++ {
		go worker(channel)
	}

	for i := 0; i <= 100; i++ {
		channel <- i
	}
}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 3)
	}
}
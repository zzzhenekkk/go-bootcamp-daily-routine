package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepSort(slice []int) chan int {
	channel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(slice))

	for _, value := range slice {
		go func(val int) {
			defer wg.Done()
			time.Sleep(time.Duration(val) * time.Millisecond)
			channel <- val
		}(value)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()
	return channel
}

func main() {
	slice := []int{100, 700, 500, 300}
	channel := sleepSort(slice)
	for val := range channel {
		fmt.Println(val)
	}
}

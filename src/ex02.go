package main

import (
	"fmt"
	"sync"
)

func multiplex(channels ...chan interface{}) chan interface{} {
	var wg sync.WaitGroup

	output := make(chan interface{})

	for _, ch := range channels {

		wg.Add(1)

		go func(c chan interface{}) {
			defer wg.Done()
			for val := range c {
				output <- val
			}
		}(ch)

	}

	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func main() {
	// Создаем три входных канала
	input1 := make(chan interface{})
	input2 := make(chan interface{})
	input3 := make(chan interface{})

	// Заполняем каналы данными
	go func() {
		for i := 0; i < 5; i++ {
			input1 <- fmt.Sprintf("input1-%d", i)
		}
		close(input1)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			input2 <- fmt.Sprintf("input2-%d", i)
		}
		close(input2)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			input3 <- fmt.Sprintf("input3-%d", i)
		}
		close(input3)
	}()

	// Используем функцию multiplex для слияния данных из трех каналов

	output := multiplex(input1, input2, input3)

	// Читаем и печатаем результаты из выходного канала
	for msg := range output {
		fmt.Println(msg)
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func crawlWeb(chanURL <-chan string) (chan string, error) {
	chanRes := make(chan string, 8)
	var wg sync.WaitGroup
	concurrentGoroutines := make(chan struct{}, 8)

	go func() {
		defer close(chanRes)
		for url := range chanURL {
			concurrentGoroutines <- struct{}{}
			wg.Add(1)

			go func(url string) {
				defer wg.Done()
				defer func() { <-concurrentGoroutines }()
				resp, err := http.Get(url)
				if err != nil {
					chanRes <- fmt.Sprintf("Error fetching URL %s: %v", url, err)
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					chanRes <- fmt.Sprintf("Error ReadAll URL %s: %v", url, err)
					return
				}

				chanRes <- string(body)
			}(url)
		}
		wg.Wait()
	}()

	return chanRes, nil
}

func main() {

	chanURL := make(chan string)

	chanRes, err := crawlWeb(chanURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		chanURL <- "https://www.simple.ru/"
		chanURL <- "http://google.com"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		chanURL <- "http://ya.ru"
		close(chanURL)
	}()

	//time.Sleep(time.Duration(6) * time.Second)
	for htmlRes := range chanRes {
		fmt.Println(htmlRes)
		fmt.Println()
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println()
	}

}

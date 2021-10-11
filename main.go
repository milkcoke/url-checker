package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Goroutines 다른함수와 동시에 실행시키는 것.
// goroutine is alive as long a program 'main function' is running
// if main function finished, goroutine should be finished
// main function doesn't wait goroutine itself

// How to communicate between main and goroutine ?
// We can send result from goroutine!
// main 은 결과를 저장하는 곳!
var successCode = [4]int{200, 201, 202, 204}
var errRequestFailed = errors.New("Request failed!")
var errInvalidRequest = errors.New("Invalid Request!")

type responseResult struct {
	url        string
	statusCode int
}

func printCount(name string) {
	for i := 0; i < 10; i++ {
		println(name, i)
		time.Sleep(time.Second)
	}
}

func waitAndCheckOdd(num int, channel chan string) {

	for i := 0; i < 10; i++ {
		println("wait for result about : ", num, i, "seconds...")
		time.Sleep(time.Second)
	}

	if num%2 == 0 {
		channel <- "false: " + strconv.Itoa(num)
	} else {
		channel <- "true " + strconv.Itoa(num)
	}
}

func isSuccess(statusCode int) bool {
	for _, value := range successCode {
		if statusCode == value {
			return true
		}
	}

	return false
}

// send only channel parameter
func hitUrl(url string, channel chan<- responseResult) error {
	//fmt.Println("전달받은 주소 : ", url)
	response, error := http.Get(url)
	// send 전용인지 receive 전용인지 명확하게 명시.

	if error != nil {
		return errRequestFailed
	}
	// send data to channel
	channel <- responseResult{url: url, statusCode: response.StatusCode}

	if isSuccess(response.StatusCode) {
		return nil
	} else {
		return errInvalidRequest
	}
}

func main() {

	channel := make(chan responseResult)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.naver.com/",
		"https://www.facebook.com/",
	}

	//waitGroup := sync.WaitGroup{}

	// goroutine!
	for index, url := range urls {
		// value_temp = range_temp[index_temp]
		// index = index_temp
		// value = value_temp

		fmt.Printf("index : %d, url: %s\n", index, url)
		go func() {
			//defer waitGroup.Done()
			// 여기로 url 이 왜 다 안넘어가지?
			fmt.Println("index in goroutine : ", index)
			err := hitUrl(url, channel)
			if err != nil {
				log.Fatalln(err)
			}
		}()
		time.Sleep(time.Millisecond * 300)
		fmt.Printf("22222 index : %d, url: %s\n", index, url)
		//go func(url string) {
		//	//defer waitGroup.Done()
		//	// 여기로 url 이 왜 다 안넘어가지?
		//	err := hitUrl(url, channel)
		//	if err != nil {
		//		log.Fatalln(err)
		//	}
		//}(url)
	}

	// 여기서는 range-based for loop 보다
	// 호출 횟수를 여러번하는게 좋음.
	// 호출 순서 상관없이 channel Message Queue (FIFO Queue) 에 먼저 추가된 메시지부터 아무거나 받아오기 때문
	for i := 0; i < len(urls); i++ {
		fmt.Printf("result: %v\n", <-channel)
	}

	// goroutine test

	/*
		// make channel
		channel := make(chan string)
		go waitAndCheckOdd(5, channel)
		go waitAndCheckOdd(2, channel)
		// When main function watch '<-' block the code (wait the message)
		// and getting a message from the channel
		// <- is receiver!
		fmt.Println("Received message from the channel : ", <-channel)
		fmt.Println("Received message from the channel : ", <-channel)

		// Go knows how many goroutines are running
		// and administrate goroutine state (message queue)
	*/
}

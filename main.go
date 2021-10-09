package main

import (
	"errors"
	"fmt"
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

func hitUrl(url string) error {
	response, error := http.Get(url)
	if error != nil {
		return errRequestFailed
	}
	if isSuccess(response.StatusCode) {
		return nil
	} else {
		return errInvalidRequest
	}
}

func main() {
	/*
		// make empty map
		results := make(map[string]string)

		urls := []string{
			"https://www.airbnb.com",
			"https://www.google.com",
			"https://www.naver.com/",
			"https://www.facebook.com/",
		}

		for _, url := range urls {
			error := hitUrl(url)
			if error != nil {
				// error type is an interface type that has an interface
				//type error interface {
				//	Error() string
				//}
				results[url] = error.Error()
			} else {
				results[url] = "Success"
			}
		}

		for url, resultString := range results {
			fmt.Println(url, " : ", resultString)
		}
	*/

	// goroutine test

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

}

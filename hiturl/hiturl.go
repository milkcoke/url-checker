package hiturl

import (
	"errors"
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

type ResponseResult struct {
	url        string
	statusCode int
}

func PrintCount(name string) {
	for i := 0; i < 10; i++ {
		println(name, i)
		time.Sleep(time.Second)
	}
}

func WaitAndCheckOdd(num int, channel chan string) {

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

func IsSuccess(statusCode int) bool {
	for _, value := range successCode {
		if statusCode == value {
			return true
		}
	}

	return false
}

// send only channel parameter
func HitUrl(url string, channel chan<- ResponseResult) error {
	//fmt.Println("전달받은 주소 : ", url)
	response, error := http.Get(url)
	// send 전용인지 receive 전용인지 명확하게 명시.

	if error != nil {
		return errRequestFailed
	}
	// send data to channel
	channel <- ResponseResult{url: url, statusCode: response.StatusCode}

	if IsSuccess(response.StatusCode) {
		return nil
	} else {
		return errInvalidRequest
	}
}

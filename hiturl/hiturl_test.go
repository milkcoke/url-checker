package hiturl

import (
	"fmt"
	"log"
	"testing"
)

func TestHitUrl(t *testing.T) {
	channel := make(chan ResponseResult)

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

		go func(url string) {
			//defer waitGroup.Done()
			// 여기로 url 이 왜 다 안넘어가지?
			err := HitUrl(url, channel)
			if err != nil {
				log.Fatalln(err)
			}
		}(url)
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

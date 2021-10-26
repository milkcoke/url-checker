package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/milkcoke/url-checker/hiturl"
	"log"
	"net/http"
	"strconv"
	"time"
)

//var successCodes = [4]int{200, 201, 202, 204}

// Go doesn't have set data-structure
// Use map[T]bool for using set space-efficiently
var successCodes = map[int]bool{
	200: true,
	201: true,
	202: true,
	204: true,
}

func isSuccessResponse(statusCode int) bool {
	_, isSuccess := successCodes[statusCode]
	fmt.Println(isSuccess)
	// status code is not 2XX
	return isSuccess
}

func isNotDigit(s string) bool {
	for _, alphabet := range s {
		if alphabet < '0' || alphabet > '9' {
			return true
		}
	}
	return false
}

func makeUrl(path string) string {
	baseUrl := "https://kr.indeed.com"
	targetUrl := baseUrl + path

	return targetUrl
}

func getPages(baseUrl string) []*hiturl.PageLink {

	// data type 헷갈림
	var pageLink = make([]*hiturl.PageLink, 0)

	response, err := http.Get(baseUrl)

	defer response.Body.Close()

	if err != nil {
		log.Fatalln(err)
	} else if isSuccessResponse(response.StatusCode) != true {
		log.Fatalf("You've got error code: %d\n", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	// 이거 동작을 어케하노
	document.Find(".pagination-list li").Each(func(i int, s *goquery.Selection) {
		// pagination class 내에 a 링크 갯수는 4개 (현재 페이지는 제외) 2, 3, 4, '>'

		pageNumberString, isExistString := s.Find("a").Attr("aria-label")
		pageUrl, isExistUrl := s.Find("a").Attr("href")

		if isNotDigit(pageNumberString) {
			return
		}

		if isExistString && isExistUrl {
			pageNumber, _ := strconv.Atoi(pageNumberString)
			pageUrl = makeUrl(pageUrl)
			// stack 에는 일단 값을 다 남기기로!
			pageLink = append(pageLink, hiturl.NewPageLink(pageNumber, pageUrl))
		}

	})

	return pageLink
}

func main() {
	var baseUrl string = "https://kr.indeed.com/jobs?q=golang&l=%EC%84%9C%EC%9A%B8&vjk=cfda1220635db9c9"
	startTime := time.Now()
	pageLinks := getPages(baseUrl)

	// index, element
	for index, pageLink := range pageLinks {
		fmt.Printf("index : %d, pageLink: %v\n", index, pageLink)
	}

	endTime := time.Now().Sub(startTime).Seconds()
	fmt.Println(endTime)

}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func main() {
	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://udemy.com",
		"https://ebay.com",
	}

	c := make(chan string)

	for _, link := range(links){
		go checkLink(link, c)
	}

	// for i := 0; i < 1000; i++{
	// 	go checkLink(<- c, c)
	// }

	// Infinite loop
	// for {
	// 	go checkLink(<- c, c)
	// }

	// Alternative syntax to loop above
	for l := range c {
		go func (link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}
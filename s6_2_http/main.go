package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// io.Copy(os.Stdout, resp.Body)

	// Interfaces do not verify your code is correct
	// They are a contract showing expectations
	io.Copy(lw, resp.Body)
}
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWtiter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWtiter{}

	io.Copy(lw, resp.Body)
}

func (logWtiter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bites:", len(bs))
	return len(bs), nil
}

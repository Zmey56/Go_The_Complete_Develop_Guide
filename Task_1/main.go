package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min_value := 1
	max_value := 10
	values := []int{}

	for i := 0; i < 10; i++ {
		values = append(values, (rand.Intn(max_value-min_value+1) + min_value))
	}

	for _, v := range values {
		if v%2 == 0 {
			fmt.Printf("%d is even\n", v)
		} else {
			fmt.Printf("%d is odd\n", v)
		}
	}
}

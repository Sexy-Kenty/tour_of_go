package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(i)
}

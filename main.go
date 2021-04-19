package main

import (
	"fmt"

	"michaelrose.dev/fb/utils"
)

func PrintAll() {
	for i := 1; i <= 100; i++ {
		fmt.Println(utils.FizzBuzz(i))
	}
}

func main() {
	PrintAll()
}

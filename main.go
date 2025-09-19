package main

import (
	"fmt"

	"education-program/input"
	"education-program/stats"
)

func main() {
	nums, err := input.ReadInts()
	if err != nil {
		fmt.Println("入力エラー:", err)
		return
	}

	fmt.Println("Sum:", stats.Sum(nums))
	// Divice
	// FizzBuzz
}

package stats

import (
	"errors"
	"strconv"
)

func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInts() ([]int, error) {
	fmt.Print("整数をスペース区切りで入力してください: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	tokens := strings.Fields(strings.TrimSpace(line))
	var nums []int
	for _, tok := range tokens {
		n, err := strconv.Atoi(tok)
		if err != nil {
			return nil, fmt.Errorf("数値でない入力があります: %v", tok)
		}
		nums = append(nums, n)
	}

	return nums, nil
}
